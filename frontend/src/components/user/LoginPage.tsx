import React, {Component} from 'react';
import {login} from '../../services/andrewwillette'
import {setBearerToken} from "../../persistence/localstorage";

export class LoginPage extends Component<any, any> {
    async sendLogin() {
        let username = (document.getElementById("username") as HTMLInputElement).value;
        let password = (document.getElementById("password") as HTMLInputElement).value;
        console.log(`calling login with ${username}, ${password}`);
        let bearerPromise = login(username, password);
        bearerPromise.then(value => {
            console.log(value.parsedBody)
            const token = value.parsedBody?.bearerToken;
            if(token){
                setBearerToken(token)
            }
        });
    }
    render() {
        return (
            <div>
                <label htmlFor={"username"}>Username</label>
                <input id={"username"} type={"text"}/>
                <label htmlFor={"password"}>Password</label>
                <input id={"password"} type={"text"}/>
                <button onClick={this.sendLogin}>Login</button>
            </div>
        );
    }
}
