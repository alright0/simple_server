import './App.css'
import {Link, Route, Routes} from "react-router-dom";
import LoginForm from "./components/login_page/login_form.jsx";

function Home() {
  return <h2>Главная страница</h2>
}

function About() {
  return <h2>О проекте</h2>
}

function NotFound() {
  return <h2>Страница не найдена</h2>
}

function App() {
  return (
      <div>
        {/* Простая навигация */}
        <nav>
          <Link to="/">Главная</Link> |{' '}
          <Link to="/about">О проекте</Link>
          <Link to="/login">Вход</Link>
        </nav>
        {/* Определяем маршруты */}
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<LoginForm />} />
          <Route path="/about" element={<About />} />
          {/* Этот маршрут сработает, если ничего не подошло */}
          <Route path="*" element={<NotFound />} />
        </Routes>
      </div>
  )
}

export default App
