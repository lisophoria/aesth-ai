"use client";

import styles from "./page.module.css";

export default function Auth() {
  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const data = new FormData(e.currentTarget);
    const email = data.get("email") as string;
    const password = data.get("password") as string;

    await fetch('/api/login', {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });
  }

  return (
    <div className={styles.flexWrapper}>
      <div className={styles.switcherWrapper}>
        <div className={styles.switcherInnerCard}>
          <p className={styles.topHeading}>Нет аккаунта?</p>
          <p className={styles.description}>Создайте его, чтобы мы могли сохранить прогресс!</p>
          <button
            className={styles.formButton}
          >
            Создать аккаунт
          </button>    
        </div>
      </div>
      <div className={styles.formContainer}>
        <div className={styles.formCard}>
        <p className={styles.title}>Войти в аккаунт</p>
          <form
            className={styles.form}
            onSubmit={handleLogin}
          >
            <input 
              type={"email"}
              className={styles.input}
              placeholder={"Email"}
              name={"email"}
              required={true}
            />
            <input
              type={"password"}
              className={styles.input}
              placeholder={"Password"}
              name={"password"}
              required={true}
            />
            <button
              className={styles.formButton}
              type={"submit"}
            >
              Войти
            </button>
          </form>
        </div>
      </div>
    </div>
  );
};