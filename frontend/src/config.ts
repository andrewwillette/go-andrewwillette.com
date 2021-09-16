export {production};

let environment : String = String(process.env.REACT_APP_ENV);
//  if (environment !== "prod" && environment !== "dev") {
//      console.error("Missing prod or dev REACT_APP_ENV declaration");
//  }
let production: boolean = environment !== "dev";
