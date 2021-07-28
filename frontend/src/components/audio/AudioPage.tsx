import React, {Component} from 'react';
import {getAllSoundcloudUrls, SoundcloudUrls} from "../../services/andrewwillette";
import {AudioPlayer} from "./AudioPlayer";

export class AudioPage extends Component<any, any> {
    generateAudioPlayers() {
        console.log("here1");
        var soundcloudUrls  = getAllSoundcloudUrls();
        return soundcloudUrls;
    }

    render() {
        var urls = this.generateAudioPlayers();
        console.log('here0');
        console.log(urls);
        urls.then((value ) => {
            var divElement = <div></div>;
            console.log("length is")
            console.log(value.parsedBody?.Urls)
            // const fields: JSX.Element[] = [];
            // for (let i = 1; i <= committedFieldsToAdd; i++) {
            //     fields.push(<Field id={i} key={i} />);
            // }
            value.parsedBody?.Urls.forEach(string => {
                console.log(string)

            })
            // console.log(value.parsedBody?.keys());
            // value.forEach(() => {
            //         console.log("for each url");
            // });
            console.log("here2");
            console.log(value);
        })
        return (
            <div>
                <AudioPlayer/>
            </div>
        );
    }
}
