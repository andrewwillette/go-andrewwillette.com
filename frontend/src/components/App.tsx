import React from 'react';
import logo from '../logo.svg';
import './App.css';
import './audio/AudioPlayer'
import {AudioPlayer} from "./audio/AudioPlayer";
import {getAllSoundcloudUrls} from "../services/andrewwillette";
import {AudioPage} from "./audio/AudioPage";



function App() {
    return (
        <div className="App">
            <p>heel</p>
            <AudioPage/>
        </div>
    );
}

export default App;
