import React from 'react';
import logo from '../logo.svg';
import './App.css';
import './audio/AudioPlayer'
import {AudioPlayer} from "./audio/AudioPlayer";

async function http<T>(
    request: RequestInfo
): Promise<T> {
    const response = await fetch(request);
    return await response.json();
}

interface BasicResponse {
    song_one: string;
    song_two: string;
}

const data = http<BasicResponse[]>(
    "http://localhost:8080"
);

console.log("here1");
console.log(data.then((todo) => {
    console.log(todo);
}));

// function getUsers() {
//     var api_response = fetch("http://localhost:8080").then((response:Response) => {
//         return response.json()
//     });
//     console.log(api_response);
// }

// async function call_local() {
//     const response = await fetch('http://localhost:8080', {
//         method: 'GET',
//         headers: {
//             'Content-Type': 'application/json'
//         }
//     });
//     console.log(response);
//     return await response.json();
// }

function App() {
    // getUsers();
    // console.log(my_value);
    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo"/>
            </header>
            <p>
                {/*{my_value}*/}
            </p>
            <AudioPlayer/>
            <AudioPlayer/>
            <AudioPlayer/>
        </div>
    );
}

export default App;
