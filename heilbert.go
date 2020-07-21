package heilbert

import (
	"errors"
)

// python implemention
// def hilbert(x, N=None, axis=-1):
// x = np.asarray(x)
// if np.iscomplexobj(x):
// 	raise ValueError("x must be real.")
// if N is None:
// 	N = x.shape[axis]
// if N <= 0:
// 	raise ValueError("N must be positive.")

// Xf = sp_fft.fft(x, N, axis=axis)
// h = np.zeros(N)
// if N % 2 == 0:
// 	h[0] = h[N // 2] = 1
// 	h[1:N // 2] = 2
// else:
// 	h[0] = 1
// 	h[1:(N + 1) // 2] = 2

// if x.ndim > 1:
// 	ind = [np.newaxis] * x.ndim
// 	ind[axis] = slice(None)
// 	h = h[tuple(ind)]
// x = sp_fft.ifft(Xf * h, axis=axis)
// return x

func Heibert(x []complex128, n int) ([]complex128, error) {
	if n == 0 {
		n = len(x)
	}

	if n <= 0 {
		return nil, errors.New("N must be positive.")
	}

	xf := FFT(x, n)
	h := make([]complex128, n)
	if n&1 == 0 {
		h[0] = 1
		h[n/2] = 1
	} else {
		h[0] = 1
		for i := 1; i <= (n+1)/2; i++ {
			h[i] = 2
		}
	}

	xx := make([]complex128, n)
	for i, v := range xf {
		xx[i] = v * h[i]
	}

	return IFFT(xx, n), nil
}
