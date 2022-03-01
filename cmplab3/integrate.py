import numpy as np
import time
import typing as t
import dataclasses as dt


@dt.dataclass
class Job:
    Method: t.Callable[[t.Callable[[float], float], float, float], float]
    Target: t.Callable[[float], float]
    Low: float
    High: float
    N: int


def fnGleb(x: float) -> float:
    return np.log(x) * np.abs(np.cos(128 * x))


def fnGlebPrecisePart(x: float) -> float:
    return (np.sin(128 * x) - 128 * x * np.cos(128 * x)) / 2097152


def fnGlebPrecise() -> float:
    high = fnGlebPrecisePart(4.141593)
    low = fnGlebPrecisePart(1)
    return high - low


def simpson(fn: t.Callable[[float], float], a: float, b: float) -> float:
    return ((b - a) / 6) * (fn(a) + 4 * fn((a + b) / 2) + fn(b))


def leftRect(fn: t.Callable[[float], float], a: float, b: float) -> float:
    return fn(a) * (b - a)


def rightRect(fn: t.Callable[[float], float], a: float, b: float) -> float:
    return fn(b) * (b - a)


def centerRect(fn: t.Callable[[float], float], a: float, b: float) -> float:
    return fn((a + b) / 2) * (b - a)


def trapezoid(fn: t.Callable[[float], float], a: float, b: float) -> float:
    return ((b - a) * (fn(a) + fn(b))) / 2


def integrate(job: Job) -> tuple[float, float]:
    h = (job.High - job.Low) / float(job.N)

    xNext: float = job.Low + h
    small: float = 0.0
    total: float = 0.0
    cnt: int = 0
    x: float = job.Low

    start = time.time()
    while xNext <= job.High:
        small = job.Method(job.Target, x, xNext)
        total += small
        x = xNext
        xNext = x + h
        cnt += 1
    took = time.time() - start

    return total, took


def pretty(methodName: str, res: float, precise: float, took: float) -> None:
    diff = res - precise
    print(
        "%s:\t%.10f\t(- pricsice = %.10f; took %0.1f seconds)"
        % (methodName, res, diff, took)
    )


if __name__ == "__main__":
    a: float = 1.0
    b: float = 21000.0
    precise = fnGlebPrecise()
    print("Precise:\t%.10f" % precise)

    nSimpson: float = 2.1e5
    res, took = integrate(
        Job(
            Method=simpson,
            Target=fnGleb,
            Low=a,
            High=b,
            N=int(nSimpson),
        ),
    )
    pretty("Simpson", res, precise, took)

    nLeft: float = 1.2e5
    res, took = integrate(
        Job(
            Method=leftRect,
            Target=fnGleb,
            Low=a,
            High=b,
            N=int(nLeft),
        ),
    )
    pretty("Left rect", res, precise, took)

    nLeft: float = 1.2e5
    res, took = integrate(
        Job(
            Method=rightRect,
            Target=fnGleb,
            Low=a,
            High=b,
            N=int(nLeft),
        ),
    )
    pretty("Right rect", res, precise, took)

    nLeft: float = 1.2e5
    res, took = integrate(
        Job(
            Method=centerRect,
            Target=fnGleb,
            Low=a,
            High=b,
            N=int(nLeft),
        ),
    )
    pretty("Center rect", res, precise, took)

    nLeft: float = 1.2e5
    res, took = integrate(
        Job(
            Method=trapezoid,
            Target=fnGleb,
            Low=a,
            High=b,
            N=int(nLeft),
        ),
    )
    pretty("Trapezoid rect", res, precise, took)
