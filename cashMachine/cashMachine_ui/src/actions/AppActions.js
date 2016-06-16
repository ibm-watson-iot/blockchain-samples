export const OPEN_SNACKBAR = "OPEN_SNACKBAR"
export const openSnackbar = () => {
  return{
    type: OPEN_SNACKBAR
  }
}

export const HIDE_SNACKBAR = "HIDE_SNACKBAR"
export const hideSnackbar = () => {
  return{
    type: HIDE_SNACKBAR
  }
}

export const SET_SNACKBAR_MSG = "SET_SNACKBAR_MSG"
export const setSnackbarMsg = (msg) => {
  return{
    type: SET_SNACKBAR_MSG,
    msg
  }
}
