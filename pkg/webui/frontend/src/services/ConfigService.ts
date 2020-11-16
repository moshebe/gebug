import axios from 'axios';


export type Config = {
  time: number;
  outgoing: boolean;
  from: string;
  to: string;
  title: string;
  body: string;
  cc: string;
  bcc: string;
};

export default class ConfigService {
    

    private static url(relative: string): string {
      const baseUrl = "http://localhost:3030" // TODO: settings
      return `${baseUrl}/${relative}`;
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
}