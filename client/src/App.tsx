import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import "./App.css";

const Header = () => {
  return (
    <>
      <div>asdfasdfasdfr</div>
      <Link to={`/about`}>asdfa</Link>
    </>
  );
};

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Header />} />
          <Route path="/about" element={<div>about</div>} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
