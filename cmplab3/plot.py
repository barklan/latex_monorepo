import matplotlib.pyplot as plt
import matplotlib as mpl
import numpy as np
import seaborn as sns


class Helper:

    def __init__(self, filename, figsize=(6, 6), axis=True):
        self.filename = filename
        self.figsize = figsize
        self.axis = axis

    def set_plot(self, figsize=(6, 6)):
        sns.set()
        sns.set_theme(style="white", palette="bright")
        COLOR = "grey"
        mpl.rcParams["text.color"] = COLOR
        mpl.rcParams["axes.labelcolor"] = COLOR
        mpl.rcParams["axes.grid"] = True
        mpl.rcParams["xtick.color"] = COLOR
        mpl.rcParams["ytick.color"] = COLOR
        mpl.rcParams.update({
            "pgf.texsystem": "pdflatex",
            'font.family': 'serif',
            'text.usetex': True,
            'pgf.rcfonts': False,
        })
        plt.figure(figsize=self.figsize)
        sns.set_context("paper")

    def save_plot(self):
        if self.axis:
            plt.axis("on")
        else:
            plt.axis("off")
        plt.savefig(self.filename, format="pgf",
                    bbox_inches="tight", pad_inches=0,
                    transparent=True)

    def __enter__(self):
        self.set_plot()
        return self

    def __exit__(self, exc_type, exc_value, exc_traceback):
        self.save_plot()


# first graphic
x = np.arange(1, np.pi+1, 0.00001)
y = 3 * x + 4 * x ** 3 - 12 * x * x - 5

y = np.log(x) * np.abs(np.cos(128*x))

with Helper("img/fig_1.pgf", figsize=(10, 10), axis=True):
    fig = plt.figure()
    fig.set_size_inches(w=8, h=5)
    ax = fig.add_subplot(111)
    ax.plot(x, y)
    # ax.plot([2.89], [0], 'o')

# second
# x = np.arange(-1.5, 1.5, 0.00001)
# y = 2 * np.tan(x) - x / 2 + 1

# with Helper("report_lab_1/images/eq_plot_2.pgf", figsize=(7, 7), axis=True):
#     fig = plt.figure()
#     fig.set_size_inches(w=3, h=2)
#     ax = fig.add_subplot(111)
#     ax.set_ylim(-8, 8)
#     ax.plot(x, y)
#     ax.plot([-0.5713], [0], 'o')
