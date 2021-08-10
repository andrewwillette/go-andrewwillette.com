import React, {Component} from 'react';
import ReactPlayer from "react-player"
// import ReactPlayer from "react-player/lazy/players/SoundCloud"

export class AudioPlayer extends Component<any, any> {
    render() {
        return (
            <div className="audioPlayer">
                <ReactPlayer
                    url = {this.props.soundcloudUrl}
                    className='react-player'
                    // config={{
                    //     // soundcloud: {
                    //     //     // should work according to docs https://github.com/CookPete/react-player but it's borked. I shove the data in url as query param lmao ??
                    //     // }
                    // }}
                />
            </div>
        );
    }
}
