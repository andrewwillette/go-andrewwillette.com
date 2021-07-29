import React, {Component} from 'react';
import {getSoundcloudUrls, SoundcloudUrl} from "../../services/andrewwillette";
import {AudioPlayer} from "./AudioPlayer";

export class AudioPage extends Component<any, any> {
    constructor(props: any) {
        super(props);
        this.state = {soundcloudUrls: []}
    }
    componentDidMount() {
        var soundcloudUrls  = getSoundcloudUrls();
        soundcloudUrls.then(soundcloudUrls => this.setState({soundcloudUrls: soundcloudUrls.parsedBody}))
    }

    renderAudioPlayers(soundcloudUrls: SoundcloudUrl[]) {
        return (
            <>
                {soundcloudUrls.map((data) => {
                    return <AudioPlayer soundcloudUrl={data.url}/>
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
