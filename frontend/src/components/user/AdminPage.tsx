import React, {Component} from 'react';
import {deleteSoundcloudUrl, getSoundcloudUrls, addSoundcloudUrl, SoundcloudUrl} from "../../services/andrewwillette";

export class AdminPage extends Component<any, any> {
    constructor(props: any) {
        super(props);
        this.state = {soundcloudUrls: []}
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
            this.updateSoundcloudUrls();
        });
    }

    addSoundcloudUrl() {
        const soundcloudUrl = (document.getElementById("addSoundCloudUrlInput") as HTMLInputElement).value;
        addSoundcloudUrl(soundcloudUrl).then(result => {
            this.updateSoundcloudUrls();
        });
    }

    renderAudioManagementList(soundcloudUrls: SoundcloudUrl[]) {
        if (soundcloudUrls === null) {
            return <></>;
        }
        return (
            <>
                {soundcloudUrls.map((data) => {
                    return (
                        <div>
                            <p>{data.url}</p>
                            <button onClick={() => this.deleteSoundcloudUrl(data.url)}>Delete URL</button>
                        </div>
                    )
                })}
            </>
        )
    }

    render() {
        const {soundcloudUrls} = this.state;
        return (
            <div>
                {this.renderAudioManagementList(soundcloudUrls)}
                <input type={"text"} id={"addSoundCloudUrlInput"}/>
                <button onClick={() => this.addSoundcloudUrl()}>Add URL</button>
            </div>
        );
    }
}
