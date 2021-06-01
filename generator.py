# Generates Golang models from the latest Redfish/Swordfish JSON schemas.
# Author: Caleb Carlson

import converter as c
import schemas
import json
import os
import shutil

# Mapping of JSON to Golang primitive types
PRIMITIVES = {
    "number": "float64",
    "integer": "int",
    "string": "string",
    "boolean": "bool"
}

# Golang types for Redfish-specific types
REDFISH_TYPES = {
    "@odata.context": "string",
    "@odata.etag": "string",
    "@odata.id": "string",
    "@odata.type": "string",
    "Id": "string",
    "Name": "string",
    "Description": "string",
    "Members": "[]map[string]interface{}",
    "Members@odata.count": "int",
    "Members@odata.nextLink": "map[string]interface{}",
    "Oem": "map[string]interface{}",
    "OemActions": "map[string]interface{}",
    "Actions": "map[string]interface{}",
    "ItemOrCollection": "map[string]interface{}",
    "ResourceCollection": "map[string]interface{}"
}

# Mapping of JSON to Golang variable names
REDFISH_VARS = {
    "@odata.context": "OdataContext",
    "@odata.etag": "OdataEtag",
    "@odata.id": "OdataId",
    "@odata.type": "OdataType",
    "Members@odata.count": "MembersOdataCount",
    "Members@odata.nextLink": "MembersOdataNextLink",
}

# Mapping of Redfish descriptions
REDFISH_DESCRIPTIONS = {
    "Oem": "The OEM extension property.",
    "OemObject": "The base type for an OEM extension.",
    "Members": "The members of this collection.",
    "Id": "The identifier that uniquely identifies the resource within the collection of similar resources.",
    "Name": "The name of the resource or array member.",
    "Description": "The description of this resource.  Used for commonality in the schema definitions."
}


def schema_name(filename):
    return filename.split(".")[0]


class Generator:

    def __init__(self, schema_dir, output_dir, name):
        self.schema_dir = schema_dir
        self.output_dir = output_dir
        self.converter = c.Converter()
        self.recreate_models_dir()
        self.name = name

    def recreate_models_dir(self):
        # Remove schemas/ dir if it already exists
        if os.path.isdir(self.output_dir):
            try:
                shutil.rmtree(self.output_dir)
            except OSError as e:
                print("Error: %s : %s" % (self.output_dir, e.strerror))

        os.mkdir(self.output_dir)
        os.mkdir(f"{self.output_dir}/models")
        os.mkdir(f"{self.output_dir}/enums")

    def generate_models(self):
        """Main entrypoint of Generator"""
        for filename in schemas.filenames(self.schema_dir):
            if filename[0].isupper():
                self.parse_file(f"{self.schema_dir}/{filename}", filename)

    def parse_file(self, file_path, filename):
        """Parse a schema .json file, generating all models/enums found within as .go files"""
        try:
            f = open(file_path, "r")
            print(f"{file_path}:")

            # Returns JSON object as a dictionary
            data = json.load(f)

            # Iterate over every model key/value pair in schema definitions list
            definitions = data["definitions"]
            for key, value in definitions.items():
                if "enum" in value:
                    print(f"  {key}: enum")
                    self.generate_enum(key, value)
                elif key[len(key) - len("Collection"):len(key)] == "Collection":  # if key ends in "Collection"
                    print(f"  {key}: collection")
                    self.generate_model(key, value["anyOf"][1])
                elif "properties" in value:
                    print(f"  {key}: model")
                    self.generate_model(key, value)

            f.close()
        except Exception as e:
            print(f"\nCaught exception in parse_file while reading file: {filename}")
            print(e)
            exit(1)

    def generate_enum(self, enum_type, enum_object):
        """Parse a schema enum object into a Golang enum"""
        enums = []  # Array of Enum objects
        enum_values = enum_object["enum"]
        enum_filename = "enum_" + self.converter.camel_to_snake(enum_type) + ".go"
        enum_dir = f"{self.output_dir}/enums"

        # Get descriptions, if they exist
        enum_descriptions = None
        has_descriptions = "enumDescriptions" in enum_object
        if has_descriptions:
            enum_descriptions = enum_object["enumDescriptions"]

        # Populate enums list
        for enum_value in enum_values:
            enum_description = None
            const_var = enum_type + "_" + self.converter.camel_to_snake(enum_value).upper()

            if has_descriptions:
                enum_description = enum_descriptions[enum_value]

            enum = Enum(enum_type, const_var, enum_value, description=enum_description)
            enums.append(enum)

        with open(f"{enum_dir}/{enum_filename}", "w") as f:

            header = f"""/* -----------------------------------------------------------------
* {enum_filename} -
*
* {self.name} {enum_type} enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */\n\n"""
            f.write(header)
            f.write("package openapi\n\n")
            f.write(f"type {enum_type} string\n\n")
            f.write("const (\n")
            for enum in enums:
                if enum.description:
                    f.write(f"\t// {enum.description}\n")
                    f.write(f"\t{enum.const_var} {enum.enum_type} = \"{enum.value}\"\n\n")
                else:
                    f.write(f"\t{enum.const_var} {enum.enum_type} = \"{enum.value}\"\n")
            f.write(")\n")

    def record_types(self):
        """Records the model types for things like Actions and user-defined objects as map[string]interface{}"""
        for filename in schemas.filenames(self.schema_dir):
            if filename[0].isupper():
                file_path = f"{self.schema_dir}/{filename}"
                try:
                    f = open(file_path, "r")
                    print(f"{file_path}:")

                    # Returns JSON object as a dictionary
                    data = json.load(f)

                    # Iterate over every model key/value pair in schema definitions list
                    definitions = data["definitions"]
                    for key, value in definitions.items():
                        if "enum" in value:
                            continue
                        elif "Collection" in key:
                            continue
                        elif "properties" in value and key not in REDFISH_TYPES:
                            properties = value["properties"]
                            if "parameters" in value or len(properties) == 0:
                                REDFISH_TYPES[key] = "map[string]interface{}"

                    f.close()
                except Exception as e:
                    print(f"\nCaught exception while reading file: {filename}")
                    print(e)
                    exit(1)

    def generate_model(self, model_type, model_value):
        """Parse a schema model object into a Golang struct"""
        if model_type in REDFISH_TYPES:
            print(f"  Predefined type. Skipping...")
            return

        try:
            model_properties = model_value["properties"]
            model_required = []
            model_description = ""
            model_filename = self.converter.camel_to_snake(model_type) + ".go"
            model_dir = f"{self.output_dir}/models"

            if "description" in model_value:
                model_description = model_value["description"]

            if "required" in model_value:
                model_required = model_value["required"]

            with open(f"{self.output_dir}/model_{model_filename}", "w") as f:
                header = f"""/* -----------------------------------------------------------------
* {model_filename} -
*
* {self.name} {model_type} resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */\n\n"""
                f.write(header)
                f.write("package openapi\n\n")

                if model_description != "":
                    f.write(f"// {model_type} - {model_description}\n")
                    f.write(f"// Modeled after {self.name} schema {model_type}\n")
                    f.write("type %s struct {\n" % model_type)

                for key, value in model_properties.items():
                    self.parse_property(f, key, value)

                f.write("}\n")

                f.close()
        except Exception as e:
            print(f"Failed while processing {model_type}")
            raise e

    def parse_property(self, file, prop_key, prop_val):
        pass

    def generate_collection(self):
        pass


class RedfishGenerator(Generator):

    def __init__(self):
        Generator.__init__(self, "schemas/redfish", "out/redfish", "DMTF Redfish")


class SwordfishGenerator(Generator):

    def __init__(self):
        Generator.__init__(self, "schemas/swordfish", "out/swordfish", "SNIA Swordfish")


class Enum:

    def __init__(self, enum_type, const_var, value, description=None):
        self.enum_type = enum_type
        self.const_var = const_var
        self.value = value
        self.description = description

    def __repr__(self):
        if self.description:
            return "Type: {}, Const val: {}, Value: {}, Description: {}".format(self.enum_type, self.const_var,
                                                                                self.value, self.description)
        else:
            return "Type: {}, Const val: {}, Value: {}".format(self.enum_type, self.const_var, self.value)
