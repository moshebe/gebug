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
  outputBinary: string
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
      const baseUrl = "http://localhost:3030" // TODO: use settings
      return `${baseUrl}/${relative}`;
    }

    static decode(config: any) : Config {
      return {
        name: config.name,
        outputBinary: config.outputBinary,
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

    static encode(config: Config): any{
      return {
        name: config.name,
        outputBinary: config.outputBinary,
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
      const res =  await axios.get(url);
      return this.decode(res.data.data.config);
    }

    static async save(location: string, config: Config) {
      const url = this.url('config');
      const payload = {
        location: location, 
        config: this.encode(config),
    };
      return axios.post(url, payload);
    }
}