import React, {Component} from 'react';

export class HomePage extends Component<any, any> {
    constructor(props: any) {
        super(props);
    }
    render() {
        return (
            <div>
                <h1>Andrew Willette</h1>
                <p>
                    I am a software and web developer currently employed at Cerner in Kansas City, Missouri, where I specialize in cloud-hosted security protocol services, both their backend implementations and frontend onboarding UIs. I like playing violin and host some recordings on my site here. Go Chiefs.
                </p>
            </div>
        );
    }
}
