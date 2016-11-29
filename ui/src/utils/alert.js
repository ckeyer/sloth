import toastr from 'toastr'

toastr.options = {
  closeButton: true,
  progressBar: true,
  showMethod: 'slideDown',
  positionClass: 'toast-top-full-width', // 'toast-top-center',
  timeOut: 4000
}

const debug = process.env.NODE_ENV !== 'production'

export default {
  success: (content, title) => {
    if (debug) {
      console.log('success', title, content)
    }
    toastr.success(content, title)
  },
  error: (content, title) => {
    if (debug) {
      console.log('error', title, content)
    }
    toastr.error(content, title)
  },
  warn: (content, title) => {
    if (debug) {
      console.log('warn', title, content)
    }
    toastr.warning(content, title)
  }
}
