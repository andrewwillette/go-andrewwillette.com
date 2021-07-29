import React, {Component} from 'react';

export class ResumePage extends Component<any, any> {
    constructor(props: any) {
        super(props);
    }
    render() {
        return (
            <div id="resume-page">

                <h3 className="resume-header">
                    Employment
                </h3>
                <ul id="resume-employment">
                    <li>
                        Cerner <text>2017-2021</text>
                        <ul>
                            <li>
                                Associate Senior Software Engineer - Cloud Identity and Access Management (Cloud IAM)
                            </li>
                        </ul>
                    </li>
                    <li>
                        GOM Software International <text>2018-2020</text>
                        <ul>
                            <li>
                                Software Engineer - Agricultural Feeding
                            </li>
                        </ul>
                    </li>
                    <li>
                        Fiddling Technologies<text>2017-2021</text>
                        <ul>
                            <li>
                                President
                            </li>
                        </ul>
                    </li>
                </ul>
                <h3 className="resume-header">
                    Experience
                </h3>
                <ul>
                    <li>
                        Java
                        <ul>
                            <li>
                                Built and extended java web services such as Cerner's enterprise user-roster, authorization-server, and single-sign-on services.
                            </li>
                            <li>
                                Common tools I use include Apache Maven, Java Spring, JAXRS, HK2, Hibernate, OracleDB, JSP, JUnit, Spock (groovy testing framework).
                            </li>
                        </ul>
                    </li>
                    <li>
                        Javascript
                        <ul>
                            <li>
                                Built React frontends for sites supporting identity-access services, such as logging into identity providers and managing user sessions.
                            </li>
                            <li>
                                Implemented front-end Typescript library for managing service calls in SSO login process, consumed by both Cloud IAM team and all Cerner web-based applications.
                            </li>
                            <li>
                                Extensive testing experience in the npm environment (mocha, jest, chai, jasmine, enzyme).
                            </li>
                            <li>
                                Experienced with C# Razor front ends consuming knockout.js for client updating.
                            </li>
                        </ul>
                    </li>
                    <li>
                        AWS
                        <ul>
                            <li>
                                Experience writing and managing Terraform configurations for different cloud environments. This includes EC2 instances, load balancers, target groups, and security groups.
                            </li>
                            <li>
                                Experience writing and managing Packer configurations for custom, per-service AMIs.
                            </li>
                            <li>

                                Deployed Jenkins master and slaves into AWS with custom jobs managed via Jenkins groovy APIs.
                            </li>
                            <li>
                                Strong familiarity with AWS billing practices.
                            </li>
                        </ul>
                    </li>
                    <li>
                        C#
                        <ul>
                            <li>
                                Built full .NET Core web application for agricultural feeding customization.
                            </li>
                            <li>
                                Implemented Braintree online payment system collecting credit and debit cards for scheduled payment processing.
                            </li>
                            <li>
                                SQL Server Schema design and implementation for feeding application management (Rations, Ration Reports, Ration Ingredients, Consultants, Producers).
                            </li>
                        </ul>
                    </li>
                    <li>
                        Miscellaneous
                        <ul>
                            <li>
                                Linux (Rhel, Ubuntu)
                            </li>
                            <li>
                                Bash, Python
                            </li>
                            <li>
                                IntelliJ Idea, Visual Studio, Vim
                            </li>
                            <li>
                                New Relic
                            </li>
                            <li>
                                GraphQL
                            </li>
                            <li>
                                Wireshark
                            </li>
                        </ul>
                    </li>
                </ul>

                <h3 className="resume-header">
                    Education: University of Iowa, 2018
                </h3>
                <ul>
                    <li>
                        Bachelor of Arts Degree - Computer Science
                    </li>
                    <li>
                        Bachelor of Arts Degree - Violin Performance
                    </li>
                </ul>
            </div>
        );
    }
}
