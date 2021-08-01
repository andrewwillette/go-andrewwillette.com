import React, {Component} from 'react';

export class LoginPage extends Component<any, any> {
    constructor(props: any) {
        super(props);
    }
    render() {
        return (
            <div>
                <label htmlFor={"username"}>Username</label>
                <input id={"username"} type={"text"}/>
                <label htmlFor={"password"}>Password</label>
                <input id={"password"} type={"text"}/>
            </div>
        );
    }
}
