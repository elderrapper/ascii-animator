# ASCII Animator

TODO:

- Add a demo video.
- Explain why project is created when the MV is ready.

## How to Use

1. Clone this repository and `cd` into it:

    ```sh
    git clone git@github.com:davidhsingyuchen/ascii-animator.git
    cd ascii-animator
    ```

1. Create `image.ans` and populate it with your favorite ASCII art (take a look at [ASCII Art Convertor](https://manytools.org/hacker-tools/convert-images-to-ascii-art/)!).

1. Create the config file:

    ```sh
    cp config.yaml.dist config.yaml
    ```

1. See how the animation looks like with the default configuration:

    ```sh
    make run
    ```

1. Tune the values in `config.yaml` according to your personal preferences.

## Notes

The stage when randomized images are displayed is designed to override the previous line on purpose.
