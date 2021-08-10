import React from 'react';
import './App.css';
import './audio/AudioPlayer'
import {AudioPage} from "./audio/AudioPage";
import {BrowserRouter, Route, Switch, Link} from "react-router-dom"
import {ResumePage} from "./resume/ResumePage";
import {HomePage} from "./home/HomePage";
import {LoginPage} from "./user/LoginPage"
import {AdminPage} from "./user/AdminPage";

function App() {
    return (
        <div className="center">
            <div className="container">
                <div className="App">
                    <BrowserRouter>
                        <h1>
                            <Link to="/">Andrew Willette</Link>
                        </h1>
                        <ul className="navbar">
                            <li>
                                <Link to="/music">Music</Link>
                            </li>
                            <li>
                                <Link to="/resume">Resume</Link>
                            </li>
                            <li>
                                <Link to="/login">Login</Link>
                            </li>
                            <li>
                                <Link to="/admin">Admin</Link>
                            </li>
                        </ul>
                        <Switch>
                            <Route exact path="/" component={HomePage}/>
                            <Route exact path="/music" component={AudioPage}/>
                            <Route path="/resume" component={ResumePage}/>
                            <Route path="/login" component={LoginPage}/>
                            <Route path="/admin" component={AdminPage}/>
                        </Switch>
                    </BrowserRouter>
                </div>
            </div>
        </div>
    );
}

export default App;
