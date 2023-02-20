import React from 'react'
import { useLocation } from 'react-router-dom'

export default function Error (): JSX.Element {
  const location = useLocation()

  return (
        <div className="flex items-center justify-center h-screen text-center">
            <div>
                <h1 className="text-6xl mb-6">
                    🃏
                </h1>
                <p>Errors happen sometimes<br/>Now is sometime 😔</p>
                <hr className={'my-2'}/>
                <a href={location.pathname} className="hover:underline">Reload</a>
            </div>
        </div>
  )
}
