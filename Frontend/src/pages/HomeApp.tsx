import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import emailLogo from "../assets/email.png";
import passwordLogo from "../assets/password.png";
import MiniRoute from "../assets/MiniRoute.png";
import cover from "../assets/security-services-Shortener-v0-1_INT.jpg";
import "./HomeApp.scss";

function HomeApp() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [rePassword, setRePassword] = useState("");
  const [loginPage, setLoginPage] = useState(true);
  const [error, setIsError] = useState("");
  const navigate = useNavigate();

  const signUp = async () => {
    try {
      if (password === rePassword) {
        const res = await axios.post("http://localhost:8080/auth/register", {
          email: email,
          password: password,
        });
        const token = res.data.Authorization;
        sessionStorage.setItem("jwt", token);
        navigate("/dashboard");
      } else {
        setIsError("Password Mismatched");
      }
    } catch (error: any) {
      if (error.response) {
        setIsError(error.response.data);
      }
    }
  };
  const signIn = async () => {
    try {
      const res = await axios.post("http://localhost:8080/auth/login", {
        email: email,
        password: password,
      });
      const token = res.data.Authorization;
      sessionStorage.setItem("jwt", token);
      navigate("/dashboard");
    } catch (error: any) {
      if (error.response) {
        setIsError(error.response.data);
      }
    }
  };
  return (
    <>
      <div className="HomeWrapper">
        <div className="AuthWrapper">
          {!loginPage ? (
            <div className="SigninandupBox">
              <h1>Sign Up</h1>
              <div>
                <img src={emailLogo} alt="" />
                <input
                  type="email"
                  placeholder="Email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div>
                <img src={passwordLogo} alt="" />

                <input
                  type="password"
                  placeholder="Password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>
              <div>
                <img src={passwordLogo} alt="" />

                <input
                  type="password"
                  placeholder="Confirm Password"
                  value={rePassword}
                  onChange={(e) => setRePassword(e.target.value)}
                />
              </div>
              <button
                onClick={(event) => {
                  event.preventDefault();
                  signUp();
                }}
                disabled={!email || !password || !rePassword}
              >
                Sign Up
              </button>
            </div>
          ) : (
            <div className="SigninandupBox">
              <h1>Sign In</h1>
              <div>
                <img src={emailLogo} alt="" />
                <input
                  type="email"
                  placeholder="Email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div>
                <img src={passwordLogo} alt="" />

                <input
                  type="password"
                  placeholder="Password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>
              <button
                onClick={(event) => {
                  event.preventDefault();
                  signIn();
                }}
                disabled={!email || !password}
              >
                Sign In
              </button>
            </div>
          )}
          <div className="CoverBox">
            <div className="logo">
              <img src={MiniRoute} alt="" />
            </div>
            <div className="cover">
              <img src={cover} alt="" />
            </div>
            <a
              onClick={() => {
                setLoginPage(!loginPage);
                setEmail("");
                setPassword("");
                setRePassword("");
              }}
            >
              {!loginPage ? "I am already member" : "Create an account"}
            </a>
          </div>
        </div>
      </div>
    </>
  );
}

export default HomeApp;
