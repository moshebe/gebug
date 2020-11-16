import axios from 'axios';


export type ConfigEnvVars = {
  envName: string
  envValue?: string
}
export type ConfigExposePort = {
  port: number
}
export type ConfigNetwork = {
  network: string
}
export type Config = {
  name: string
  outputBinPath: string
  buildCommand: string
  runCommand: string
  runtimeImage: string
  debuggerEnabled: boolean
  debuggerPort: number
  networks: ConfigNetwork[]
  exposePorts: ConfigExposePort[]
  environment: ConfigEnvVars[]  
};

export default class ConfigService {
    
  private static url(relative: string): string {
      const baseUrl = "http://localhost:3030" // TODO: settings
      return `${baseUrl}/${relative}`;
    }

    static decodeModel(config: any) : Config {
      return {
        name: config.name,
        outputBinPath: config.outputBinPath,
        buildCommand:  config.buildCommand,
        runCommand:  config.runCommand,
        runtimeImage: config.runtimeImage,
        debuggerEnabled: config.debuggerEnabled,
        debuggerPort: config.debuggerPort ? parseInt(config.debuggerPort) : config.debuggerPort,
        exposePorts: config.exposePorts.map(p => ({port: p})),
        networks: config.networks.map(n => ({network: n})),
        environment: config.environment.filter(e => e).map(e => {
          const parts = e.split('=');
          if (parts.length > 1)
            return ({envName: parts[0], envValue: parts[1]});
          else
            return ({envName: parts[0]});
        })
      };
    }

    static encodeModel(config: Config): any{
      return {
        name: config.name,
        outputBinPath: config.outputBinPath,
        buildCommand:  config.buildCommand,
        runCommand:  config.runCommand,
        runtimeImage: config.runtimeImage,
        debuggerEnabled: config.debuggerEnabled === true,
        debuggerPort: Number(config.debuggerPort),
        exposePorts: config.exposePorts.filter(i => i).map(p => p.port),
        networks: config.networks.filter(i => i).map(n => n.network),
        environment: config.environment.filter(i => i).map(e => e.envName + (e.envValue ? `=${e.envValue}` : ''))
      };
    }

    static async get(path: string) {
      const url = this.url(`config?path=${path}`); // TODO: input sanitization
      const x =  await axios.get(url);      
      // console.log('x: ', x);
      // console.log('x1: ', x.data);
      console.log('x2: ', x.data.data.config);
      return x;
      // .then(
          // response => {                      
            // this.config = response.data.data.config;
            // console.log('got response from sever, set config to: ', this.config);
        // }
    }

    static async save(config: Config) {
      const url = 'https://enlpbbatnnqf.x.pipedream.net'; // TODO: replace with the backend endpoint
      await axios.post(url, config);
    }
}