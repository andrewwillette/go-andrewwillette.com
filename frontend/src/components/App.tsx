import React from 'react';
import './App.css';
import './audio/AudioPlayer'
import {AudioPage} from "./audio/AudioPage";
import {BrowserRouter, Route, Switch } from "react-router-dom"
import {ResumePage} from "./resume/ResumePage";

function App() {
    return (
        <div className="App">
            <BrowserRouter>
                <Switch>
                    <Route exact path="/" component={AudioPage} />
                    <Route path="/resume" component={ResumePage} />
                </Switch>
            </BrowserRouter>
        </div>
    );
}

export default App;
