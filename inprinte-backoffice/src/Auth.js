
export const httpClient = () => {
    const { token } = JSON.parse(localStorage.getItem('auth')) || {};
    return { Authorization: `Bearer ${token}` };
  };
  
  export const authProvider = {
    // authentication
    login: ({ username, password }) => {
      const request = new Request(
        'http://localhost:8080/api/login',
        {
          method: 'POST',
          body: JSON.stringify({ email: username, password  }),
        }
      );
      return fetch(request)
        .then((response) => {
          if (response.status < 200 || response.status >= 300) {
            throw new Error(response.statusText);
          }
          return response.json();
        })
        .then((auth) => {
          localStorage.setItem(
            'auth',
            JSON.stringify({ ...auth, fullName: username })
          );
        })
        .catch(() => {
          throw new Error('Network error');
        });
    },
    checkError: (error) => {
      const status = error.status;
      if (status === 401 || status === 403) {
        localStorage.removeItem('auth');
        return Promise.reject();
      }
      // other error code (404, 500, etc): no need to log out
      return Promise.resolve();
    },
    checkAuth: () =>
      localStorage.getItem('auth')
        ? Promise.resolve()
        : Promise.reject({ message: 'login required' }),
    logout: () => {
      localStorage.removeItem('auth');
      document.cookie = "role=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/";
      return Promise.resolve();
    },
    getIdentity: () => {
      try {
        const { id, fullName, avatar } = JSON.parse(localStorage.getItem('auth'));
        return Promise.resolve({ id, fullName, avatar });
      } catch (error) {
        return Promise.reject(error);
      }
    },
    getPermissions: (params) => Promise.resolve(),
  };