import React, { type KeyboardEvent, useMemo, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { ArrowRightIcon, PlusIcon } from '@heroicons/react/24/solid'

const API_HOST = process.env.REACT_APP_API_HOST as string

export default function Home (): JSX.Element {
  const navigate = useNavigate()
  const [roomQuery, setRoomQuery] = useState<string>('')

  const validRoom = useMemo(() => {
    return roomQuery.length === 36
  }, [roomQuery])

  const createLobby = (): void => {
    fetch(`${API_HOST}/lobby/create`, {
      credentials: 'include'
    }).then((res) => {
      if (res.status === 201) {
        res.json().then((json) => {
          navigate(`/lobby/${json.lobby_id as string}`)
        }).catch((e) => {
          console.error(e)
        })
      } else {
        console.error(res)
      }
    }).catch((e) => {
      console.error(e)
    })
  }

  const goToRoom = (): void => {
    navigate('/lobby/' + roomQuery)
  }

  const handleKeyDown = (event: KeyboardEvent): void => {
    if (event.key === 'Enter' && validRoom) {
      goToRoom()
    }
  }

  return <div className="isolate bg-white">
        <main>
            <div className="relative px-6 lg:px-8">
                <div className="mx-auto max-w-3xl pt-20 pb-32 sm:pt-48 sm:pb-40">
                    <div>
                        <div>
                            <h1 className="text-4xl font-bold tracking-tight text-center sm:text-6xl">
                                🃏 Planning Poker
                            </h1>
                            <p className="mt-6 muted-text text-center">
                                Join or create a room and play planning poker in realtime.
                            </p>
                            <div className="mt-12 sm:flex block gap-x-4 sm:justify-center flex-wrap">
                                <input
                                    className="appearance-none border-2 sm:w-96 w-full h-16 sm:h-auto text-center focus:border-green-500 rounded-lg focus:ring-0 focus:ring-offset-0"
                                    type="text"
                                    autoCapitalize={'off'}
                                    autoComplete={'off'}
                                    value={roomQuery}
                                    onChange={(e) => { setRoomQuery(e.target.value) }}
                                    onKeyDown={handleKeyDown}
                                    placeholder={'your lobby id here'}
                                />
                                <button
                                    onClick={() => { goToRoom() }}
                                    disabled={!validRoom}
                                    className="sm:w-auto w-full mt-4 h-16 sm:h-auto sm:mt-0 flex items-center justify-center rounded-lg bg-gray-600 disabled:cursor-not-allowed px-4 py-1.5 text-base font-semibold leading-7 text-white shadow-sm ring-1 ring-gray-600 hover:bg-gray-700 hover:ring-gray-700">
                                    <span>Join</span>
                                    <ArrowRightIcon className="ml-2 h-5 w-5"/>
                                </button>
                            </div>
                            <div className="mt-4 sm:flex block gap-x-4 sm:justify-center flex-wrap items-center">
                                <span className="hidden sm:block muted-text">Or</span>
                                <button
                                    onClick={() => { createLobby() }}
                                    className="sm:w-auto w-full mt-4 h-16 sm:h-auto sm:mt-0 flex items-center justify-center rounded-lg bg-green-600 px-4 py-1.5 text-base font-semibold leading-7 text-white shadow-sm ring-1 ring-green-600 hover:bg-green-700 hover:ring-green-700">
                                    <span>Create</span>
                                    <PlusIcon className="ml-2 h-5 w-5"/>
                                </button>
                                <span className="hidden sm:block muted-text">a new lobby</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div className={'font-medium text-center text-gray-400'}>
                Created by <a target="_blank" rel="noreferrer" href="https://github.com/frederikhs" className="hover:underline">frederikhs</a>.
                View on <a target="_blank" rel="noreferrer" href="https://github.com/frederikhs/planning-poker" className="hover:underline">GitHub</a>
            </div>
        </main>
    </div>
}
