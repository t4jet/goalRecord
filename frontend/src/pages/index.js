import { useEffect, useState } from 'react';

export default function Home() {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/api/tasks', {
      mode: 'cors',
    })
      .then((response) => response.json())
      .then((data) => setTasks(data));
  }, []);

  return (
    <div>
      <h1>Tasks</h1>
      <ul>
        {tasks.map((task) => (
          <li key={task.ID}>
            {task.Name} - {task.Finished ? 'Finished' : 'Not Finished'}
          </li>
        ))}
      </ul>
    </div>
  );
}

