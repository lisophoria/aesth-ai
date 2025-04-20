import styles from "./page.module.css";

export default function Home() {
  return (
    <div className={styles.page}>
      <main className={styles.main}>
        <h1>Aesth AI</h1>

        <ol>
          <li>
            Amazing description of the project.
          </li>
          <li>
            And maybe some steps to start using it.
          </li>
        </ol>

        <div className={styles.controls}>
          <a
            className={styles.primary}
            target="_blank"
            rel="noopener noreferrer"
          >
            Try it now
          </a>
          <a
            target="_blank"
            rel="noopener noreferrer"
            className={styles.secondary}
          >
            How it works?
          </a>
        </div>
      </main>
      <footer className={styles.footer}>
        <a
          href="https://github.com/aesth-ai"
          target="_blank"
          rel="noopener noreferrer"
        >
          GitHub
        </a>
      </footer>
    </div>
  );
}
