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
    "Members": "map[string]string",
    "Members@odata.count": "int",
    "Members@odata.nextLink": "map[string]string",
    "Oem": "map[string]interface{}",
    "OemActions": "map[string]interface{}",
    "Actions": "map[string]interface{}",
    "ItemOrCollection": "map[string]interface{}",
    "ResourceCollection": "map[string]interface{}",
    "count": "int",
    "etag": "string",
    "context": "string",
    "id": "string",
    "type": "string",
    "nextLink": "map[string]string"
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


def type_of_items(items):
    """Processes the "items" object, which comes in 3 flavors:
        1. "items": {
                    "type": [
                        "string",
                        "null"
                    ]
                },

        2. "items": {
                    "anyOf": [
                        {
                            "$ref": "#/definitions/RoleMapping"
                        },
                        {
                            "type": "null"
                        }
                    ]
                },

        3. "items": {
                    "$ref": "http://redfish.dmtf.org/schemas/v1/Zone.json#/definitions/Zone"
                },
    """
    if "anyOf" in items:
        json_type = items["anyOf"][0]["$ref"].split("/")[-1]
    elif "$ref" in items:
        json_type = items["$ref"].split("/")[-1]
    else:
        json_type = items["type"][0]

    if json_type in REDFISH_TYPES:
        return REDFISH_TYPES[json_type]
    elif json_type in PRIMITIVES:
        return PRIMITIVES[json_type]
    else:
        return json_type


def parse_property(f, prop_key, prop_val, is_required):
    """Parses a JSON schema property into a Golang model field, with a comment."""
    try:

        # Capture the property description if it exists, and add it as a comment.
        description = ""
        if prop_key in REDFISH_DESCRIPTIONS:
            description = REDFISH_DESCRIPTIONS[prop_key]
        elif "description" in prop_val:
            description = prop_val["description"]

        if description != "":
            f.write(f"\t// {description}\n")

        # Use the property key as the variable name, fixing it as needed for @odata keys which aren't valid var names.
        variable = prop_key
        if prop_key in REDFISH_VARS:
            variable = REDFISH_VARS[prop_key]

        # Some variables have @odata.count on them
        if "@odata.count" in variable:
            variable = variable.split("@")[0] + "OdataCount"

        golang_type = determine_prop_type(prop_val)

        print(f"Golang type for {prop_key}: {golang_type}")
        if is_required:
            f.write("\t%s %s `json:\"%s\"`\n\n" % (variable, golang_type, prop_key))
        else:
            f.write("\t%s %s `json:\"%s,omitempty\"`\n\n" % (variable, golang_type, prop_key))

    except Exception as e:
        print(f"Failed while processing property {prop_key}")
        raise e


def determine_prop_type(prop_val):
    """Determines the appropriate Golang type of the property. This is the crux of model generation.
        A.) There are two formats the JSON property values come in. The first includes a "type" subfield, which makes it
            much easier to determine the type:

            The "type" field comes in several flavors:
            1.) "type": <json primitive>, (i.e. "type": "string", "type": "number", "type": "boolean")
                This is the simplest case, and uses a JSON primitive -> Golang primitive mapping to determine the type.
            2.) "type": [<type 1>, <type 2>, ..., <type n>], (i.e. "type": ["string", "null"])
                This is the second simplest case, and in all cases seen it is a list of non-null prim type followed by
                a null-type.
            3.) "type": "array"
                This means the type should be an array of values. It also implies there is an "items" array which lists
                the type of the items in the array, which can be any of cases A.1), A.2), or B.)

        B.) The other doesn't include a type field but instead has a "$ref" field, which may or may not be embedded into
            an "anyOf" array of objects. This looks like:
            "anyOf": [
                {
                    "$ref": "#/definitions/MaxPrefix"
                },
                {
                    "type": "null"
                }
            ]
            Thus, we always choose the first "$ref" type.
    """
    golang_type = ""
    json_type = ""

    # Case A.)
    if "type" in prop_val:

        # Case A.1), A.3)
        if type(prop_val["type"]) is str:
            json_type = prop_val["type"]
        # Case A.2)
        else:
            # Iterate over types and choose the first non-"null" type
            for v_type in prop_val["type"]:
                if v_type != "null":
                    json_type = v_type

        if json_type in PRIMITIVES:  # map JSON -> Golang primitive
            golang_type = PRIMITIVES[json_type]
        elif json_type in REDFISH_TYPES:  # map to predefined Golang type, i.e. Oem -> map[string]interface{}
            golang_type = REDFISH_TYPES[json_type]
        elif json_type == "array":  # recursive call to determine type of items val
            golang_type = "[]" + determine_prop_type(prop_val["items"])
        elif json_type == "object":  # user-defined object type
            golang_type = "map[string]interface{}"

    # Case B.)
    else:
        reference = ""  # Reference URI to Redfish/Swordfish definition
        if "anyOf" in prop_val:
            reference = prop_val["anyOf"][0]["$ref"]
        elif "$ref" in prop_val:
            reference = prop_val["$ref"]

        # Parse reference URI
        ending = reference.split("/")[-1]
        if ending in REDFISH_TYPES:  # map to predefined Golang type, i.e. Oem -> map[string]interface{}
            golang_type = REDFISH_TYPES[ending]
        elif "autoExpand" in prop_val or "http" not in reference:  # Golang type should be reference type, not link
            golang_type = ending
        else:  # Golang type should be a link to type, i.e. { "@odata.id": "/redfish/v1/StorageServices/12345" }
            golang_type = "map[string]string"

    return golang_type


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
                    if "anyOf" not in value and "properties in value":
                        print(f"  {key}: model")
                        self.generate_model(key, value, filename)
                    else:
                        print(f"  {key}: collection")
                        self.generate_model(key, value["anyOf"][1], filename)
                elif "properties" in value:
                    print(f"  {key}: model")
                    self.generate_model(key, value, filename)

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
                            REDFISH_TYPES[key] = key
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

    def generate_model(self, model_type, model_value, filename):
        """Parse a schema model object into a Golang struct"""
        if model_type in REDFISH_TYPES:
            print(f"  Predefined type. Skipping...")
            return

        try:
            model_properties = model_value["properties"]
            model_required = []
            model_description = ""
            model_filename = "model_" + self.converter.camel_to_snake(model_type) + ".go"
            model_dir = f"{self.output_dir}/models"

            if "description" in model_value:
                model_description = model_value["description"]

            if "required" in model_value:
                model_required = model_value["required"]

            with open(f"{model_dir}/{model_filename}", "w") as f:
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
                    if self.name == "DMTF Redfish":
                        f.write(f"// Modeled after {self.name} schema https://redfish.dmtf.org/schemas/v1/{filename}\n")
                    else:
                        f.write(f"// Modeled after {self.name} schema https://redfish.dmtf.org/schemas/swordfish/{filename}\n")
                    f.write("type %s struct {\n" % model_type)

                for key, value in model_properties.items():
                    parse_property(f, key, value, key in model_required)

                f.write("}\n")

                f.close()
        except Exception as e:
            print(f"Failed while processing {model_type}")
            raise e


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
