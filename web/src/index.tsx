import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { createBrowserRouter, Link, RouterProvider } from 'react-router-dom'
import Lobby from './Pages/Lobby'
import Home from './Pages/Home'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Home/>
  },
  {
    path: 'lobby/:lobbyId',
    element: <Lobby/>
  },
  {
    path: '*',
    element: <p>404 <Link className={'text-blue-600 dark:text-blue-500 hover:underline'} to={'/'}><br/>Go home</Link></p>
  }
])

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
)

root.render(
    // <React.StrictMode>
    <RouterProvider router={router}/>
    // </React.StrictMode>
)
