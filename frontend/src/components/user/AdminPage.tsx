import React, {Component} from 'react';
import {deleteSoundcloudUrl, getSoundcloudUrls, addSoundcloudUrl, SoundcloudUrl, login} from "../../services/andrewwillette";
import {setBearerToken} from "../../persistence/localstorage";
import {UnauthorizedBanner} from "./UnauthorizedBanner";
import {LoginSuccessBanner} from "./LoginSuccessBanner";

export class AdminPage extends Component<any, any> {
    constructor(props: any) {
        super(props);
        this.state = {soundcloudUrls: [], unauthorizedReason: null, loginSuccess: false}

        this.sendLogin = this.sendLogin.bind(this);
    }

    componentDidMount() {
        this.updateSoundcloudUrls();
    }

    updateSoundcloudUrls() {
        getSoundcloudUrls().then(soundcloudUrls => {
            this.setState({soundcloudUrls: soundcloudUrls.parsedBody});
        });
    }

    deleteSoundcloudUrl(soundcloudUrl: string) {
        deleteSoundcloudUrl(soundcloudUrl).then(result => {
            if(result.status === 201 || result.status === 200) {
                this.setState({unauthorizedReason: null});
            } else {
                this.setState({unauthorizedReason: "Not logged in, cannot delete URLS"});
            }
            this.updateSoundcloudUrls();
        });
    }

    addSoundcloudUrl() {
        const soundcloudUrl = (document.getElementById("addSoundCloudUrlInput") as HTMLInputElement).value;
        addSoundcloudUrl(soundcloudUrl).then(result => {
            // if result is 401, show unauthorized banner
            this.updateSoundcloudUrls();
        });
    }

    async sendLogin() {
        let username = (document.getElementById("username") as HTMLInputElement).value
        let password = (document.getElementById("password") as HTMLInputElement).value

        let responsePromise = login(username, password)
        responsePromise.then(response => {
            if(response.status === 200) {
                const token = response.parsedBody
                if(token) {
                    setBearerToken(String(token))
                    this.setState({unauthorizedReason: null, loginSuccess: true})
                }
            } else {
                this.setState({unauthorizedReason: "Login Failed", loginSuccess: false});
            }
        });
    }

    renderAdminBanner(unauthorizedReason: string, loginSuccess: boolean) {
        if (unauthorizedReason !== null) {
            return <UnauthorizedBanner unauthorizedReason={unauthorizedReason}/>
        } else if (loginSuccess) {
            return <LoginSuccessBanner/>
        } else {
            return <></>
        }
    }

    renderAudioManagementList(soundcloudUrls: SoundcloudUrl[]) {
        if (soundcloudUrls === null) {
            return <></>;
        }
        return (
            <>
                {soundcloudUrls.map((data) => {
                    return (
                        <div key={data.url}>
                            <p>{data.url}</p>
                            <button key={data.url} onClick={() => this.deleteSoundcloudUrl(data.url)}>Delete URL</button>
                        </div>
                    )
                })}
            </>
        )
    }

    render() {
        const {soundcloudUrls, unauthorizedReason, loginSuccess} = this.state;
        return (
            <div>
                <div>
                    {this.renderAdminBanner(unauthorizedReason, loginSuccess)}
                </div>
                <div>
                    <label htmlFor={"username"}>Username</label>
                    <input id={"username"} type={"text"}/>
                    <label htmlFor={"password"}>Password</label>
                    <input id={"password"} type={"text"}/>
                    <button onClick={this.sendLogin}>Login</button>
                </div>
                {this.renderAudioManagementList(soundcloudUrls)}
                <input type={"text"} id={"addSoundCloudUrlInput"}/>
                <button onClick={() => this.addSoundcloudUrl()}>Add URL</button>
            </div>
        );
    }
}
