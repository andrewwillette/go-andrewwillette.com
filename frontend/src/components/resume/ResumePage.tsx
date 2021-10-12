import React, {Component} from 'react';
import "./resume.css";

export class ResumePage extends Component<any, any> {
    render() {
        return (
            <div id="resume-page">
                <h3 className="resume-header">
                    Software Development Experience
                </h3>
                <ul>
                    <li>
                        Cerner - Software Engineer II - Cloud Identity and Access Management - 2017 - 2021
                        <ul>
                            <li>
                                Added support in enterprise User Roster for inactive users, allowing for persistence of old account resources rather than deleting them.
                            </li>
                            <li>
                                Added enterprise support for JSON Web Token(JWT) authentication. Configured application-registered accounts to allow registration of JSON Web Key Sets and Authorization Server validation of associated signed JWTs.
                            </li>
                            <li>
                                Updated enterprise System Accounts solution to allow for automatic registration of lower-privileged third-party accounts. Part of compliance with 21st Century Cures Act regulations for allowing greater interoperability of EHR APIs.
                            </li>
                            <li>
                                Implemented New Relic and Google Analytics client libraries in many different Tomcat-hosted cloud solutions. Both Counters and Gauges configured for UI interaction and website performance.
                            </li>
                            <li>
                                Built React frontends from wireframes for sites supporting identity-access services, such as logging into identity providers and managing user sessions.
                            </li>
                            <li>
                                Wrote javascript client code for consuming Cloud IAM Graph-QL endpoints which facilitated mapping clients to their identity-providers.
                            </li>
                            <li>
                                Implemented npm Typescript client library used by internal and third-party apps for managing authentication with Cerner Single-Sign-On solution.
                            </li>
                            <li>
                                Extensive CSS experience styling login pages, user tables, and account customization UIs.
                            </li>
                            <li>
                                Managed linux server deployments with Chef Infrastructure Management. Good understanding of centralized Chef server and per-service cookbook/role configurations.
                            </li>
                            <li>
                                Wrote and managed Terraform scripts for EC2 instances, Elastic Load Balancers, Target Groups, and Security Groups in a multi-cloud-region deployment pattern.
                            </li>
                            <li>
                                Experience writing and managing Packer configurations for per-service AMIs.
                            </li>
                            <li>
                                Integrated AWS environment with Jenkins master / slaves deployments. Using native jenkins groovy API's and AMI's managed via Packer, all jenkins builds operated in AWS infrastructure.
                            </li>
                            <li>
                                Extensive testing experience in the npm environment (mocha, jest, chai, jasmine, enzyme) and java environment (failsafe and surefire).
                            </li>
                            <li>
                                Primary Java technologies used include Maven, Spring, JAXRS, HK2, Hibernate, Tomcat, OracleDB, JSP, JUnit, and Spock.
                            </li>
                        </ul>
                    </li>
                    <li>
                        GOM Software International - Software Engineer - Agricultural Feeding, 2018 - 2020
                        <ul>
                            <li>
                                Full stack .NET Core 2.0 development using razor frontends.
                            </li>
                            <li>
                                Implemented eight different frontend pages from wireframe specs, complete with full table functionality managed via knockout.js.
                            </li>
                            <li>
                                Implemented Braintree online payment system collecting credit and debit cards for a scheduled payment processing plan.
                            </li>
                            <li>
                                Extended Microsoft SQL Server schema for managing livestock feeding data.
                            </li>
                            <li>
                                Implemented role-based authentication.
                            </li>
                        </ul>
                    </li>
                    <li>
                        Fiddling Technologies - President, 2017 - 2021
                        <ul>
                            <li>
                                Personal website built on Golang rest API server and Typescript React Single-Page-Application front-end.
                            </li>
                            <li>
                                Managed amateur violin performances around Kansas City, primarily coffee shops and busking.
                            </li>
                        </ul>
                    </li>
                    <li>
                        Miscellaneous
                        <ul>
                            <li>
                                Editor workflows: Vim, IntelliJ Idea.
                            </li>
                            <li>
                                Strong unix knowledge with Linux (Ubuntu, RHEL) and MacOS.
                            </li>
                            <li>
                                Additional development experience with Bash, Golang, Python, Ruby, and C++.
                            </li>
                            <li>
                                Moderate containerization experience with Docker.
                            </li>
                            <li>
                                Strong knowledge configuring applications with New Relic and Splunk for metrics and troubleshooting.
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
