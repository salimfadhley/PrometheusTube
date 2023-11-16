import { NavLink, useSearchParams } from "@remix-run/react";
import { BellAlertIcon } from "@heroicons/react/24/outline";
import Avatar from "@mui/material/Avatar";
import Login from "app/components/login";
import Register from "app/components/register";

import Modal from "@mui/material/Modal";
import React from "react";
import { UserState } from "~/state";
import { useEffect, useState } from "react";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import ChangePassword from "./changepassword";

export function Navbar({ displayAvatar }) {
  const [isHydrated, setIsHydrated] = useState(false);
  useEffect(() => {
    setIsHydrated(true);
  }, []);

  const [open, setOpen] = React.useState(false);
  const [showAvatar, setShowAvatar] = React.useState(displayAvatar);

  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  const [openLogin, setOpenLogin] = React.useState(false);
  const handleOpenLogin = () => setOpenLogin(true);
  const handleCloseLogin = () => setOpenLogin(false);

  const [openReset, setOpenReset] = React.useState(false);
  const handleOpenReset = () => setOpenReset(true);
  const handleCloseReset = () => setOpenReset(false);

  let [searchParams, setSearchParams] = useSearchParams();
  let loggedIn = UserState((state) => state.loggedIn);
  let uid = UserState((state) => state.UID);
  let admin = UserState((state) => state.admin);


  // menu nav stuff
  const [anchorEl, setAnchorEl] = React.useState(null);
  const openMenu = Boolean(anchorEl);
  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleCloseMenu = () => {
    setAnchorEl(null);
  };

  let setLoggedIn = UserState((state) => state.setLoggedIn);
  let setUID = UserState((state) => state.setUID);

  function handleLogout() {
    setLoggedIn(false);
    setUID(null); // TODO clear the jwt cookie
  }

  return (
    <div className="flex justify-between bg-white-50 w-screen py-2 px-6">
      <div className="col-start-1 pt-1 self-start w-36">
        <NavLink
          className="text-cherry-red-100 font-extrabold text-text-single-400"
          to="/"
        >
          PrometheusTube
        </NavLink>
      </div>
      <div className="inline-block w-96">
        <input
          placeholder="Search"
          value={searchParams.get("search") ?? ""}
          className="w-full bg-white-300 rounded-full w-full pl-3 p-1"
          onChange={async (event) => {
            setSearchParams((prev) => {
              prev.set("search", event.currentTarget.value);
              return prev;
            });
          }}
        />
      </div>
      <div className="inline-block text-right w-36 ">
        {!loggedIn || !isHydrated ? (
          <span className="float-left">
            <span>
              <button
                onClick={handleOpenLogin}
                className="rounded-full py-1 px-2 border-cherry-red-200 text-cherry-red-200  border-2	"
              >
                <span className="float-left inline-block relative">Login</span>
              </button>
            </span>
            <span>
              <button
                onClick={handleOpen}
                className="rounded-full py-1 px-2 text-cherry-red-100	"
              >
                <span className="float-left inline-block relative">
                  Register
                </span>
              </button>
            </span>
          </span>
        ) : (
          <span>
            <span className="w-5 inline-block relative align-middle">
              <BellAlertIcon className="w-5 text-cherry-red-100"></BellAlertIcon>
            </span>
            <span
              onClick={handleClick}
              className="w-5 inline-block relative align-middle ml-1"
            >
              <Avatar sx={{ width: 32, height: 32 }}>{uid}</Avatar>
            </span>
            <Menu
              id="basic-menu"
              anchorEl={anchorEl}
              open={openMenu}
              onClose={handleCloseMenu}
              MenuListProps={{
                "aria-labelledby": "basic-button",
              }}
            >
              <NavLink to={"/profile/" + uid}>
                <MenuItem>Profile</MenuItem>
              </NavLink>
              <MenuItem onClick={handleOpenReset}>Reset Password</MenuItem>
              {admin ? <NavLink to={"/archive-requests/"}>
                <MenuItem>Archive requests</MenuItem>
              </NavLink>
              : null
              }
              <MenuItem onClick={handleLogout}>Logout</MenuItem>
            </Menu>
          </span>
        )}
      </div>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Register></Register>
      </Modal>
      <Modal
        open={openLogin}
        onClose={handleCloseLogin}
        aria-labelledby="modal-modal-login"
        aria-describedby="modal-modal-login"
      >
        <Login
          setLogin={() => {
            setOpenLogin();
            setShowAvatar(true);
          }}
        ></Login>
      </Modal>
      <Modal
        open={openReset}
        onClose={handleCloseReset}
        aria-labelledby="modal-modal-reset"
        aria-describedby="modal-modal-reset"
      >
        <ChangePassword handleClose={handleCloseReset}></ChangePassword>
      </Modal>
    </div>
  );
}
