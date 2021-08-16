import React, {Component} from 'react';
import {getSoundcloudUrls, SoundcloudUrl} from "../../services/andrewwillette";
import {AudioPlayer} from "./AudioPlayer";
import "./audio.css";

export class AudioPage extends Component<any, any> {
    constructor(props: any) {
        super(props);
        this.state = {soundcloudUrls: []}
    }
    componentDidMount() {
        getSoundcloudUrls().then(soundcloudUrls => {
            this.setState({soundcloudUrls: soundcloudUrls.parsedBody})
        });
    }

    renderAudioPlayers(soundcloudUrls: SoundcloudUrl[]) {
        if(soundcloudUrls === null) {
            return <></>;
        }
        return (
            <>
                {soundcloudUrls.map((data) => {
                    return <AudioPlayer key={data.url} soundcloudUrl={data.url}/>
                })}
            </>
        )
    }
    render() {
        const {soundcloudUrls} = this.state;
        return (
            <div>
                {this.renderAudioPlayers(soundcloudUrls)}
            </div>
        );
    }
}
