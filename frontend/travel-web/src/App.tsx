import React from "react";
import { BrowserRouter } from "react-router-dom";
import "react-toastify/dist/ReactToastify.css";

import "./assets/style/App.scss";
import "./assets/style/_custome.scss";
import "bootstrap/dist/css/bootstrap.min.css";

import AppRoutes from "./router/AppRoutes";
import { ToastContainer } from "react-toastify";
import ScrollToTop from "./components/ScrollToTop";

const App: React.FC = () => {
  return (
    <>
      <BrowserRouter>
        <ScrollToTop />

        <AppRoutes />

        <ToastContainer
          position="top-right"
          autoClose={1200}
          hideProgressBar={false}
          newestOnTop={false}
          closeOnClick
          pauseOnFocusLoss
          pauseOnHover={false}
        />
      </BrowserRouter>
    </>
  );
};

export default App;
