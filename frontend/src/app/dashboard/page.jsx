'use client'

import { useEffect, useState } from "react";

export default function Dashboard() {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(`/api/dashboard/`, {
          method: "GET",
        });
        const responseJSON = await response.json();
        const setTasks = responseJSON.data;
      } catch (error) {
        console.error(error)
      }
      fetchData();
    }
  }, [])
  return (
    <div>
      <ul>
        {tasks.map((task) => (
        <li key={task.id}>{task.name}</li>
        ))}
      </ul>
    </div>
  )
}

