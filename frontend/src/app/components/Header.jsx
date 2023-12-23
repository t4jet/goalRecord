'use client'
import styles from '../styles/app.module.scss';
import Button from './Button';

function Header() {

  return (
    <div className={styles.appHeader}>
      <h1>全てのタスク</h1>
      <Button variant="primary" onClick={() => setModalOpen(true)}>
        Add Task
      </Button>
    </div>
    
  );
}

export default Header;
