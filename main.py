import schemas
import generator as g


# Update these with latest .zip urls as needed. The latest urls can be found at
DMTF_SCHEMA_ZIP_URL = "https://www.dmtf.org/sites/default/files/standards/documents/DSP8010_2021.1.zip"
SNIA_SCHEMA_ZIP_URL = "https://www.snia.org/sites/default/files/swordfish/draft/v1.2.2/zip/Swordfish_v1.2.2_Schema.zip"


def main():
    # Download and prune schemas
    schemas.recreate_schemas_dir()
    schemas.download_redfish_schemas(DMTF_SCHEMA_ZIP_URL)
    schemas.download_swordfish_schemas(SNIA_SCHEMA_ZIP_URL)
    schemas.prune()

    rf_generator = g.RedfishGenerator()
    sf_generator = g.SwordfishGenerator()
    rf_generator.record_types()
    sf_generator.record_types()
    rf_generator.generate_models()
    sf_generator.generate_models()


if __name__ == '__main__':
    main()
