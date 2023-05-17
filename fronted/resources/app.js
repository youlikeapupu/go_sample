// import React from "react";
import ReactDOM from "react-dom";
import axios from 'axios';
import React, { useState } from 'react';

function  App() {
    let number = 0
    const [number, setNumber] = useState(0)

    const getSquare = ()=> {
        axios.get('/controller/square', {
            params: {
                value:  9,
            }
        })
        .then((res) => {
            console.log('res.data.data: ' + res.data.data)
            // res.data.data: 81

            number = res.data.data
            console.log('number: ' + number)
            // number: 81
            // 但畫面上的 number 仍為 0
        })
        .catch((err) => {
            console.log(err);
        });
    }

    return (
        <>
            <div>
                原本數字：
                <input value="9"></input>
                <button onClick={ ()=> setNumber(99) }>取得平方</button>
            </div>
            <br/>
            <h1>{ number }</h1>
        </>
    )
}

ReactDOM.render(<App />, document.getElementById('app'));
