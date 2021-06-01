# Downloads the latest Redfish/Swordfish JSON schemas.
# Author: Caleb Carlson

import wget
import os
from os import walk
import zipfile
import shutil


def recreate_schemas_dir():
    # Remove schemas/ dir if it already exists
    if os.path.isdir("schemas"):
        try:
            shutil.rmtree("schemas")
        except OSError as e:
            print("Error: %s : %s" % ("schemas", e.strerror))

    os.mkdir("schemas")
    os.mkdir("schemas/redfish")
    os.mkdir("schemas/swordfish")


# Returns a list of directory names in a directory
def dir_names(directory):
    d = []
    for (_, d_names, _) in walk(directory):
        d.extend(d_names)
        break
    return d


# Returns a list of filenames in a directory
def filenames(directory):
    f = []
    for (_, _, f_names) in walk(directory):
        f.extend(f_names)
        break
    return f


# Find latest "Redfish Schema" .zip url on https://www.dmtf.org/dsp/DSP8010
def download_redfish_schemas(zip_url):
    directory = "schemas/redfish/"
    zip_file = f"{directory}/redfish.zip"
    wget.download(zip_url, zip_file)
    try:
        with zipfile.ZipFile(zip_file) as z:
            z.extractall(directory)
            print("Extracted all")
    except Exception as e:
        print("Invalid file: " + str(e))

    dsp_name = dir_names(directory)[0]
    for filename in filenames(f"{directory}/{dsp_name}/json-schema"):
        shutil.move(f"{directory}/{dsp_name}/json-schema/{filename}", f"{directory}/{filename}")

    os.remove(zip_file)
    if os.path.isdir(f"{directory}/{dsp_name}"):
        try:
            shutil.rmtree(f"{directory}/{dsp_name}")
        except OSError as e:
            print("Error: %s : %s" % (f"{directory}/{dsp_name}", e.strerror))


# Find the latest "Swordfish Schema and Registries Bundle" .zip url on https://www.snia.org/forums/smi/swordfish
def download_swordfish_schemas(zip_url):
    directory = "schemas/swordfish/"
    zip_file = f"{directory}/swordfish.zip"
    wget.download(zip_url, zip_file)
    try:
        with zipfile.ZipFile(zip_file) as z:
            z.extractall(directory)
            print("Extracted all")
    except Exception as e:
        print("Invalid file: " + str(e))

    for dir_name in dir_names(directory):
        if dir_name != "json-schema":
            try:
                shutil.rmtree(f"{directory}/{dir_name}")
            except OSError as e:
                print("Error: %s : %s" % (f"{directory}/{dir_name}", e.strerror))

    os.remove(zip_file)
    for filename in filenames(f"{directory}/json-schema"):
        shutil.move(f"{directory}/json-schema/{filename}", f"{directory}/{filename}")

    os.removedirs(f"{directory}/json-schema")


# Deletes all but the latest versions of the schemas
def prune_not_latest():
    resources = {}
    redfish_filenames = filenames("schemas/redfish")
    swordfish_filenames = filenames("schemas/swordfish")

    for filename in redfish_filenames:
        parts = filename.split(".")
        resources[parts[0]] = ".".join(parts)  # Replace dictionary entry with latest version

    for filename in swordfish_filenames:
        parts = filename.split(".")
        resources[parts[0]] = ".".join(parts)  # Replace dictionary entry with latest version

    for filename in redfish_filenames:
        if filename not in resources.values():
            os.remove(f"schemas/redfish/{filename}")

    for filename in swordfish_filenames:
        if filename not in resources.values():
            os.remove(f"schemas/swordfish/{filename}")