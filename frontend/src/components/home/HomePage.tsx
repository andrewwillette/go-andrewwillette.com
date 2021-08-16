import React, {Component} from 'react';
import homepage_photo from './homepage_photo.jpg';

export class HomePage extends Component<any, any> {
    render() {
        return (
            <>
                <img src={homepage_photo} className="personalImage" alt="logo" />
                <div>
                    <p id={"personalBio"}>
                        I am a software developer based in Kansas City, Kansas, where I work primarily with cloud-hosted security protocol services, both their backend implementations and frontend on-boarding UIs. I like playing violin and host some recordings on my site here.
                    </p>
                </div>
            </>
        );
    }
}
