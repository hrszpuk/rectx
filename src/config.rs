use std::fs;
use std::process::exit;
use serde_derive::{Deserialize, Serialize};


#[derive(Debug, Deserialize, Serialize)]
struct Config {
    pub project_name: String,
}

impl Config {
    pub fn new(name: String) -> Config {
        Config {
            project_name: name,
        }
    }

    pub fn generate(&self, path: String) -> std::io::Result<()> {
        let toml_data = match toml::to_string(&self) {
            Ok(data) => data,
            Err(error) => {  // Temporary hacky fix -- definitely change later
                println!("ERROR! {:?}", error);
                exit(1);
                String::new()
            }
        };
        fs::write(path + "/config.toml", toml_data)?;
        Ok(())
    }

    /*pub fn load(path: String) -> std::io::Result<Config> {
        let data = fs::read(path)?;
        toml::from_slice(&*data)?
    }*/
}