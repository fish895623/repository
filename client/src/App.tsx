import { Link, Route, Routes } from "react-router-dom";
import Page from "./DayReport";
import "./App.css";

const Header = () => {
  return (
    <>
      <div>asdfasdfasdfr</div>
      <Link to="/report">asdfa</Link>
    </>
  );
};

const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Header />} />
      <Route path="/report" element={<Page />} />;
    </Routes>
  );
};

export default App;
