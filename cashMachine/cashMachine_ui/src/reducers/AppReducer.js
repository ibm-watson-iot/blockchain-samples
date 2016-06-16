import * as AppActions from '../actions/AppActions'

//a reducer to store global state
export const app = (state={
  ui:{
    snackbar:{
      open: false,
      msg: ""
    }
  }
}, action) =>{
  switch(action.type){
    case AppActions.OPEN_SNACKBAR:
      return Object.assign({}, state, {
        ui:{
          ...state.ui,
          snackbar:{
            ...state.ui.snackbar,
            open:true
          }
        }
      })
    case AppActions.HIDE_SNACKBAR:
      return Object.assign({}, state, {
        ui:{
          ...state.ui,
          snackbar:{
            ...state.ui.snackbar,
            open:false
          }
        }
      })
    case AppActions.SET_SNACKBAR_MSG:
      return Object.assign({}, state, {
        ui:{
          ...state.ui,
          snackbar:{
            ...state.ui.snackbar,
            msg:action.msg
          }
        }
      })
    default:
      return state;
  }
}
