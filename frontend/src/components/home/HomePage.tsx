import React, {Component} from 'react';
import homepage_photo from './homepage_photo.jpg';
import "./homePage.css";

export class HomePage extends Component<any, any> {
    render() {
        return (
            <>
                <img src={homepage_photo} className="personalImage" alt="logo" />
                <div id="home-page">
                    <p id={"personalBio"}>
                        Hi! My name is Andrew Willette. I am a software developer based in Kansas City, Kansas, where I work primarily with cloud-hosted security protocol services, both their backend implementations and frontend management UIs. I like playing violin and host some recordings on my site here.
                    </p>
                </div>
            </>
        );
    }
}
