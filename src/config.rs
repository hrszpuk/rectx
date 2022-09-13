/**  Config module
*   This module manages the config.toml inside of a ReCTx project!
*   Features:
*   - Generating new config.toml files (for the manager module)
*   - Reading config.toml files into the Config struct
*   - Processing any issues regarding the config.toml (such as missing or incorrect fields)
 **/

use std::fs;
use std::process::exit;
use serde_derive::{Deserialize, Serialize};


/// Config structure will store all the contents within
/// the config.toml so Manager can use it turning compilation.
#[derive(Debug, Deserialize, Serialize)]
struct Config {
    pub project_name: String,
}

impl Config {

    /// Simple constructor for Config struct
    pub fn new(name: String) -> Config {
        Config {
            project_name: name,
        }
    }

    /// Generates a config.toml file
    pub fn generate(&self, path: String) -> std::io::Result<()> {

        // TODO: This function needs to be tested fyi
        let toml_data = match toml::to_string(&self) {
            Ok(data) => data,
            Err(error) => {  // Temporary hacky fix -- definitely change later
                println!("ERROR! {:?}", error);
                exit(1);  // TODO: Add CLI error handlers
                String::new()
            }
        };

        // Finally, we create the config.toml with our config data
        fs::write(path + "/config.toml", toml_data)?;
        Ok(())
    }

    /*pub fn load(path: String) -> std::io::Result<Config> {
        let data = fs::read(path)?;
        toml::from_slice(&*data)?
    }*/
}