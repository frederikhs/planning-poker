import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import {createBrowserRouter, Link, RouterProvider,} from "react-router-dom";
// import Home from "./Home";
import Lobby from "./Lobby";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Lobby/>,
    },
    // {
    //     path: "lobby/:lobby_id",
    //     element: <Lobby/>,
    // },
    {
        path: "*",
        element: <p className={"muted-text"}>Hmm. This does not exist. Visit <Link className={"text-blue-600 dark:text-blue-500 hover:underline"} to={"/"}>the
            homepage</Link></p>
    }
]);

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);

root.render(
    // <React.StrictMode>
    <RouterProvider router={router}/>
    // </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
