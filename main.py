import json
import os
from os import walk
import converter as c
import schemas

# Update these with latest .zip urls as needed
DMTF_SCHEMA_ZIP_URL = "https://www.dmtf.org/sites/default/files/standards/documents/DSP8010_2021.1.zip"
SNIA_SCHEMA_ZIP_URL = "https://www.snia.org/sites/default/files/swordfish/draft/v1.2.2/zip/Swordfish_v1.2.2_Schema.zip"

# Golang types for JSON
json_primitives = {
    "number": "float64",
    "integer": "int",
    "string": "string",
    "boolean": "bool"
}

# Golang types for Redfish
redfish_types = {
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
    "Actions": "map[string]interface{}"
}

redfish_vars = {
    "@odata.context": "OdataContext",
    "@odata.etag": "OdataEtag",
    "@odata.id": "OdataId",
    "@odata.type": "OdataType",
    "Members@odata.count": "MembersOdataCount",
    "Members@odata.nextLink": "MembersOdataNextLink",
}

redfish_descriptions = {
    "Oem": "The OEM extension property.",
    "OemObject": "The base type for an OEM extension.",
    "Members": "The members of this collection.",
    "Id": "The identifier that uniquely identifies the resource within the collection of similar resources.",
    "Name": "The name of the resource or array member.",
    "Description": "The description of this resource.  Used for commonality in the schema definitions."
}

enums_output_dirs = "/home/inf0rmatiker/go/src/test/models"  # "out/enums"
models_output_dirs = "/home/inf0rmatiker/go/src/test/models"  # "out/models"
collections_output_dirs = "/home/inf0rmatiker/go/src/test/models"  # "out/collections"
Converter = c.Converter()


class Enum:

    def __init__(self, enum_type, const_var, value, description):
        self.enum_type = enum_type
        self.const_var = const_var
        self.value = value
        self.description = description

    def __repr__(self):
        return "Type: {}, Const val: {}, Value: {}, Description: {}".format(self.enum_type, self.const_var, self.value,
                                                                            self.description)


class Collection:

    def __init__(self):
        pass


class Model:

    def __init__(self):
        pass


def generate_enums(enum_type, items):
    enums = []  # Array of Enum objects
    enum_descriptions = items["enumDescriptions"]
    enum_values = items["enum"]
    enum_filename = "enum_" + Converter.camel_to_snake(enum_type) + ".go"

    for enum_value in enum_values:
        enum_description = enum_descriptions[enum_value]
        const_var = enum_type + "_" + Converter.camel_to_snake(enum_value).upper()
        enums.append(Enum(enum_type, const_var, enum_value, enum_description))

    with open(f"{enums_output_dirs}/{enum_filename}", "w") as f:

        header = f"""/* -----------------------------------------------------------------
 * {enum_filename} -
 *
 * DMTF Redfish {enum_type} enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */\n\n"""
        f.write(header)
        f.write("package openapi\n\n")
        f.write(f"type {enum_type} string\n\n")
        f.write("const (\n")
        for enum in enums:
            f.write(f"\t// {enum.description}\n")
            f.write(f"\t{enum.const_var} {enum.enum_type} = \"{enum.value}\"\n\n")
        f.write(")\n")


def generate_collection(key, item):
    print(f"  {key}: collection")
    obj = item["anyOf"][1]
    collection_description = obj["description"]
    properties = obj["properties"]
    required_props = obj["required"]
    rsrc = key[:-len("Collection")]
    collection_filename = "collection_" + Converter.camel_to_snake(rsrc) + ".go"
    with open(f"{collections_output_dirs}/{collection_filename}", "w") as f:
        header = f"""/* -----------------------------------------------------------------
* {collection_filename} -
*
* DMTF Redfish {key} resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */\n\n"""
        f.write(header)
        f.write("package openapi\n\n")
        f.write(f"// {key} - {collection_description}\n")
        f.write("type %s struct {\n" % key)
        for key, value in properties.items():

            description = ""
            if key in redfish_descriptions:
                description = redfish_descriptions[key]
            elif "description" in value:
                description = value["description"]

            if description != "":
                f.write(f"\t// {description}\n")

            variable = key
            if key in redfish_vars:
                variable = redfish_vars[key]

            type = ""
            if key in redfish_types:
                type = redfish_types[key]

            is_required = key in required_props
            if is_required:
                f.write("\t%s %s `json:\"%s\"`\n\n" % (variable, type, key))
            else:
                f.write("\t%s %s `json:\"%s,omitempty\"`\n\n" % (variable, type, key))

        f.write("}\n")


def generate_model(key, item, name):
    print(f"  {key}: model")

    # Skip Actions for now
    if key == "Actions" or "parameters" in item:
        return

    try:
        model_description = ""
        if "description" in item:
            model_description = item["description"]

        model_properties = item["properties"]

        model_required = []
        if "required" in item:
            model_required = item["required"]

        model_filename = Converter.camel_to_snake(key) + ".go"

        with open(f"{models_output_dirs}/model_{model_filename}", "w") as f:
            header = f"""/* -----------------------------------------------------------------
* {model_filename} -
*
* DMTF Redfish {key} resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */\n\n"""
            f.write(header)
            f.write("package openapi\n\n")

            if model_description != "":
                f.write(f"// {key} - {model_description}\n")
            f.write(f"// Modeled after Redfish schema {name}\n")
            f.write("type %s struct {\n" % key)

            for key, value in model_properties.items():
                description = ""
                if key in redfish_descriptions:
                    description = redfish_descriptions[key]
                elif "description" in value:
                    description = value["description"]

                if description != "":
                    f.write(f"\t// {description}\n")

                variable = key
                if key in redfish_vars:
                    variable = redfish_vars[key]
                if "@odata.count" in variable:
                    variable = variable.split("@")[0] + "OdataCount"

                json_type = ""
                if "type" in value:
                    if type(value["type"]) is str:
                        json_type = value["type"]
                        if json_type in json_primitives:
                            json_type = json_primitives[json_type]
                    else:
                        for v_type in value["type"]:
                            if v_type != "null":
                                json_type = json_primitives[v_type]
                                break

                if json_type == "" and key in redfish_types:
                    json_type = redfish_types[key]

                # Custom object type, need to get reference
                if json_type == "":
                    reference = ""
                    if "anyOf" in value:
                        reference = value["anyOf"][0]["$ref"]
                    elif "$ref" in value:
                        reference = value["$ref"]

                    json_type = reference.split("/")[-1]
                    # if "autoExpand" in value:
                    #     if value["autoExpand"]:
                    #         json_type = reference.split("/")[-1]
                    # elif "http" not in reference:
                    #     json_type = reference.split("/")[-1]
                    # else:
                    #     json_type = "map[string]string"


                if json_type == "array":
                    items = value["items"]
                    if "type" in items:
                        item_types = items["type"]
                        if type(item_types) is str:
                            json_type = f"[]{item_types}"
                        else:
                            json_type = f"[]{item_types[0]}"

                    elif "$ref" in items:
                        item_types = items["$ref"].split("/")[-1]
                        if item_types == "idRef":
                            json_type = "[]map[string]string"
                        else:
                            json_type = f"[]{item_types}"

                    elif "anyOf" in items:
                        item_types = items["anyOf"][0]["$ref"].split("/")[-1]
                        if item_types == "idRef":
                            json_type = "[]map[string]string"
                        else:
                            json_type = f"[]{item_types}"

                if json_type == "count":
                    json_type = "int"

                is_required = key in model_required
                if is_required:
                    f.write("\t%s %s `json:\"%s\"`\n\n" % (variable, json_type, key))
                else:
                    f.write("\t%s %s `json:\"%s,omitempty\"`\n\n" % (variable, json_type, key))

            f.write("}\n")

    except Exception as e:
        print(f"Failed while processing {key}")
        raise e


def generate_typedef(key, item, name):
    item_type = json_primitives[item["type"]]


def parse_filename(filename):
    if "Collection.json" in filename:
        return filename[:-5], True
    else:
        parts = filename.split(".")
        return f"{parts[0]}.json", False


def parse_resource(filepath):
    try:
        f = open(filepath, "r")
        print(f"{filepath}:")

        # returns JSON object as a dictionary
        data = json.load(f)

        items_to_skip = [
            "Description",
            "Id",
            "Name",
            "Oem",
            "Identifier",
            "Item",
            "ItemOrCollection",
            "Links",
            "Location",
            "OemObject",
            "ReferenceableMember",
            "Resource",
            "ResourceCollection"
        ]

        definitions = data["definitions"]
        for key in definitions.keys():
            item = definitions[key]
            if key not in items_to_skip:
                if "type" in item:
                    if "properties" in item and item["type"] == "object":
                        generate_model(key, item, "Resource.json")
                    else:
                        generate_typedef(key, item, "Resource.json")
        f.close()
    except Exception as e:
        pass


def parse_file(filepath, filename):
    name, is_collection = parse_filename(filename)
    if filename == "Resource.json":
        parse_resource(filepath)
        return

    try:
        f = open(filepath, "r")
        print(f"{filepath}:")

        # returns JSON object as a dictionary
        data = json.load(f)

        definitions = data["definitions"]
        for key in definitions.keys():
            item = definitions[key]
            if "enum" in item:
                print(f"  {key}: enum")
                generate_enums(key, item)
            elif is_collection:
                generate_collection(key, item)
            elif "type" in item:
                if "properties" in item and item["type"] == "object":
                    generate_model(key, item, filename)
                else:
                    generate_typedef(key, item, filename)

        f.close()

    except Exception as e:
        print(f"\nCaught exception while reading file: {filename}")
        print(e)
        exit(1)


def main():
    schemas.recreate_schemas_dir()
    schemas.download_redfish_schemas(DMTF_SCHEMA_ZIP_URL)
    schemas.download_swordfish_schemas(SNIA_SCHEMA_ZIP_URL)
    schemas.prune_not_latest()

    schema_prefix = "rfsf/Redfish"
    files_to_skip = [
        "odata.v4_0_5.json",
        "odata-v4.json",
        "redfish-error.v1_0_1.json",
        "redfish-schema-v1.json",
        "redfish-schema.v1_8_0.json",
        "redfish-payload-annotations.v1_1_1.json"
    ]

    # f = []
    # for (_, _, filenames) in walk(schema_prefix):
    #     f.extend(filenames)
    #     break
    #
    # for filename in f:
    #     if filename not in files_to_skip:
    #         parse_file(f"{schema_prefix}/{filename}", filename)




if __name__ == '__main__':
    main()
