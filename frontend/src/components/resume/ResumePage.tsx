import React, {Component} from 'react';
import "./resume.css";

export class ResumePage extends Component<any, any> {
    render() {
        return (
            <div id="resume-page">

                <h3 className="resume-header">
                    Employment
                </h3>
                <ul id="resume-employment">
                    <li>
                        Cerner 2017 - 2021
                        <ul>
                            <li>
                                Software Engineer II - Cloud Identity and Access Management
                            </li>
                        </ul>
                    </li>
                    <li>
                        GOM Software International 2018 - 2020
                        <ul>
                            <li>
                                Software Engineer - Agricultural Feeding
                            </li>
                        </ul>
                    </li>
                    <li>
                        Fiddling Technologies 2017 - 2021
                        <ul>
                            <li>
                                President
                            </li>
                        </ul>
                    </li>
                </ul>
                <h3 className="resume-header">
                    Software Development Experience
                </h3>
                <ul>
                    <li>
                        Java
                        <ul>
                            <li>
                                Added support in Cerner's User Roster backend and frontend services for inactive users, rather than deleting resources.
                            </li>
                            <li>
                                Added support in Cerner's Authorization Server for JSON Web Key Set authentication. Includes the work to allow registering JSON Web Key Sets to pair with the signed JWT's presented to the Authorization Server.
                            </li>
                            <li>
                                Updated Cerner's System Accounts application to allow for lower-privileged third-party System Accounts automatic registration. Part of 21st Century Cures Act for allowing greater interoperability of Cerner's APIs.
                            </li>
                            <li>
                                Implemented New Relic and Google Analytics Metrics client library in many different Tomcat-hosted cloud solutions. Both Counters and Gauges configured for UI interaction and website performance.
                            </li>
                            <li>
                                Primary Java technologies include Maven, Spring, JAXRS, HK2, Hibernate, Tomcat, OracleDB, JSP, JUnit, and Spock.
                            </li>
                        </ul>
                    </li>
                    <li>
                        Javascript / Typescript
                        <ul>
                            <li>
                                Built React frontends from wireframes for sites supporting identity-access services, such as logging into identity providers and managing user sessions.
                            </li>
                            <li>
                                Wrote javascript consuming Cloud IAM Graph-QL endpoints for mapping clients to their identity-providers.
                            </li>
                            <li>
                                Implemented npm Typescript client library for third-party apps to authenticate with Cerner Single-Sign-On sessions.
                            </li>
                            <li>
                                Extensive css experience styling login pages, user tables, and account customization UIs.
                            </li>
                            <li>
                                Extensive testing experience in the npm environment (mocha, jest, chai, jasmine, enzyme).
                            </li>
                            <li>
                                Personal website built on Typescript-based React front end.
                            </li>
                        </ul>
                    </li>
                    <li>
                        AWS
                        <ul>
                            <li>
                                Experience writing and managing Terraform scripts for EC2 instances, Elastic Load Balancers, Target Groups, and Security Groups in a multi-cloud-region deployment pattern.
                            </li>
                            <li>
                                Experience writing and managing Packer configurations for per-service AMIs.
                            </li>
                            <li>
                                Integrated AWS environment with Jenkins master / slaves deployments. Using native jenkins groovy API's and AMI's managed via Packer, all jenkins builds operated in AWS infrastructure.
                            </li>
                            <li>
                                Allowed passivity in cloud migrations via customized Elastic Load Balancers rules to route legacy traffic to newly supported resource locations.
                            </li>
                            <li>
                                Strong familiarity with AWS billing practices / patterns.
                            </li>
                        </ul>
                    </li>
                    <li>
                        C#
                        <ul>
                            <li>
                                Implemented eight different frontend pages from wireframe specs, complete with full table functionality managed via knockout.js.
                            </li>
                            <li>
                                Implemented Braintree online payment system collecting credit and debit cards for a scheduled payment processing plan.
                            </li>
                            <li>
                                Managed Microsoft SQL Server schema modifications for role-based authentication.
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
                                Bash, Golang, Python
                            </li>
                            <li>
                                IntelliJ Idea, Visual Studio, Vim
                            </li>
                            <li>
                                Docker
                            </li>
                            <li>
                                New Relic
                            </li>
                            <li>
                                GraphQL
                            </li>
                            <li>
                                Chef CI/CD Infrastructure
                            </li>
                        </ul>
                    </li>
                </ul>

                <h3 className="resume-header">
                    Education: University of Iowa, 2018
                </h3>
                <ul>
                    <li>
                        Bachelor's Degree - Computer Science
                    </li>
                    <li>
                        Bachelor's Degree - Violin
                    </li>
                </ul>
            </div>
        );
    }
}
