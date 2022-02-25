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


def fnIlyana(x: float) -> float:
    return 1 / (x * (3 * x + 2) - 1)


def ilyanaPecise() -> float:
    return np.log(3.0) / 4.0


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
    precise = ilyanaPecise()
    print("Precise:\t%.10f" % precise)

    nSimpson: float = 2.1e5
    res, took = integrate(
        Job(
            Method=simpson,
            Target=fnIlyana,
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
            Target=fnIlyana,
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
            Target=fnIlyana,
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
            Target=fnIlyana,
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
            Target=fnIlyana,
            Low=a,
            High=b,
            N=int(nLeft),
        ),
    )
    pretty("Trapezoid rect", res, precise, took)
