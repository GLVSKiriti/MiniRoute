import MiniRoute from "../assets/MiniRoute.png";
import { useNavigate } from "react-router-dom";
import "./NavBar.scss";

function NavBar() {
  const navigate = useNavigate();
  const logout = () => {
    sessionStorage.removeItem("jwt");
    navigate("/");
  };

  return (
    <div className="navBar">
      <div className="logo">
        <img src={MiniRoute} alt="" />
      </div>
      <div className="menu">
        <div className="myurls">My Urls</div>
        <div className="logout" onClick={() => logout()}>
          LogOut
        </div>
      </div>
    </div>
  );
}

export default NavBar;