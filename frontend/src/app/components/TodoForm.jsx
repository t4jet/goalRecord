'use client'
import { useState } from 'react';
import style from '../styles/app.module.scss';
import Button from './Button';

export default function TodoForm() {
  const [data, setData] = useState('')

  const onChangeHandler = (e) => {
    setData(e.target.value)
  }

  const onSubmitHandler = async (e) => {
    e.preventDefault()

    const res = await fetch('/api/post', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ data }),
    })

    const data = await res.json()
    setPostedData(data.body)
  }
  return (
    <form onSubmit={onSubmitHandler} className={style.appHeader} action='/api/form' method='POST'>
      <input value={data} className={style.emptyText} onChange={onChangeHandler} type='text' name='data' placeholder='タスクを入力' />
      <Button variant="primary" type='onSubmit'>
        追加
      </Button>
    </form>
    
  );
}

