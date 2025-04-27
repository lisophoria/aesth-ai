"use client";

import styles from "@/app/common.module.css";
import { Fragment } from "react";

export default function Auth() {
  // const [isRegisterMode, setRegisterMode] = useState<boolean>(false);

  const auth = async () => {
    console.log('calling auth');
    await fetch('/api/login', {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email: "lisophoria@mail.com" }),
    });
  }

  const ping = async () => {
    await fetch('/api/ping', {
      method: "GET",
      credentials: "include",
    })
  }

  return (
    <div className={styles.page}>
      <div className={styles.main}>
        <Fragment>
          <h1>Logging in</h1>
          <div className={styles.controls}>
            <a
              onClick={() => auth()}
              className={styles.primary}
            >
              Log In
            </a>
            <a 
              onClick={() => ping()}
              className={styles.secondary}>
              Send ping to server
            </a>
          </div>
        </Fragment>
      </div>
    </div>
  )
};