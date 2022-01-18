import React, {Component} from 'react';
import "./resume.css";

export class ResumePage extends Component<any, any> {
    render() {
        return (
            <div id="resume-page">
                <h1 className="resume-header">
                    Andrew Willette - Software Engineer
                </h1>
                <text className="leftFloat">Links</text>
                <ul>
                    <li>
                        Github: https://github.com/andrewwillette
                    </li>
                    <li>
                        Linkedin: https://www.linkedin.com/in/andrew-willette-28792811a/
                    </li>
                </ul>
                <h2 className="resume-header">
                    Software Development Experience
                </h2>
                Cerner - Software Engineer II - Cloud Identity and Access Management - 2017 - 2022
                <ul>
                    <li>
                        Added support in enterprise User Roster for inactive users, allowing for persistence of old account resources rather than deleting them.
                    </li>
                    <li>
                        Added Service Account support for JSON Web Token(JWT) authentication. Configured application-registered accounts to allow registration of JSON Web Key Sets and Authorization Server validation of associated signed JWTs.
                    </li>
                    <li>
                        Write and manage Terraform configuration files for EC2 instances, Elastic Load Balancers, Target Groups, and Security Groups in a multi-cloud-region deployment pattern.
                    </li>
                    <li>
                        Updated enterprise Service Accounts to allow for automatic registration of lower-privileged third-party accounts. Part of compliance with "21st Century Cures Act" federal regulations for allowing greater interoperability of EHR APIs.
                    </li>
                    <li>
                        Implemented New Relic and Google Analytics client libraries in many different Tomcat-hosted cloud services. Both Counters and Gauges configured for UI interaction and website performance.
                    </li>
                    <li>
                        Built React frontends from wireframes for sites supporting identity-access services, such as logging into identity providers and managing user sessions.
                    </li>
                    <li>
                        Wrote javascript client code for consuming Cloud IAM Graph-QL endpoints which facilitated mapping tenants/clients to their identity-providers.
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
                        Write and manage Packer configurations for per-service AMIs.
                    </li>
                    <li>
                        Integrated AWS environment with Jenkins master / slaves deployments. Using native jenkins groovy API's and AMI's managed via Packer, all jenkins builds operated in AWS infrastructure.
                    </li>
                    <li>
                        Extensive testing experience in the npm environment (mocha, jest, chai, jasmine, enzyme) and java environment (failsafe and surefire).
                    </li>
                    <li>
                        Primary Java technologies used include Spring, JAXRS, HK2, Hibernate, Tomcat, Maven, OracleDB, JSP, JUnit, and Spock.
                    </li>
                </ul>
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
                Fiddling Technologies - President, 2017 - 2022
                <ul>
                    <li>
                        Personal website built on Golang rest API server and Typescript React Single-Page-Application front-end.
                    </li>
                    <li>
                        Managed amateur violin performances around Kansas City, primarily coffee shops and busking.
                    </li>
                </ul>
                Miscellaneous
                <ul>
                    <li>
                        Editor workflows: IntelliJ Idea, Neovim.
                    </li>
                    <li>
                        Strong unix knowledge with Linux (Ubuntu, RHEL) and MacOS.
                    </li>
                    <li>
                        Additional development experience with Golang, Bash, Python, Ruby, and C++.
                    </li>
                    <li>
                        Moderate containerization experience with Docker.
                    </li>
                </ul>
                <h2 className="resume-header">
                    Education: University of Iowa, 2018
                </h2>
                <li>
                    Bachelor's Degree - Computer Science
                </li>
                <li>
                    Bachelor's Degree - Violin
                </li>
            </div>
        );
    }
}
