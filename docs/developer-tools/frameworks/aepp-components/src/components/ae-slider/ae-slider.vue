<template>
  <agile class="ae-slider" :class="{ fullscreen }" :options="defaults">
    <slot />
  </agile>
</template>
<script>
import { Agile } from 'vue-agile';

/**
 * The ae-slider uses vue-agile under the hood: https://github.com/lukaszflorczak/vue-agile
 *
 * The option is the configuration object, where you can modify
 * the behavior of the vue-agile.
 *
 * Currently this is a simple wrapper on top of the vue-agile
 * with some stylistic changes, and acts like a interface
 * in case in the future we will have to change / implement our
 * own slider.
 */
export default {
  name: 'ae-slider',
  components: { Agile },
  props: {
    /**
     * You can modify the behavior of the slider
     * using the following options:
     * @link https://github.com/lukaszflorczak/vue-agile#options
     */
    options: Object,

    /**
     * Creates a full screen slider, where the element tries to go
     * 100% height and 100% width of the parent element, and places
     * the round dots inside of the slider
     */
    fullscreen: Boolean,
  },
  computed: {
    /**
     * Returns the default object
     * configuration for vue-agile
     * @return {{}}
     */
    defaults() {
      return Object.assign({
        arrows: false,
        infinite: false,
      }, this.options);
    },
  },
};
</script>
<style lang="scss" scoped>
@import '../../styles/globals/functions';
@import '../../styles/variables/animations';
@import '../../styles/variables/colors';

.ae-slider {
  position: relative;
  width: 100%;
}

.ae-slider /deep/ .agile__slide {
  padding: rem(20px) rem(32px);
}

.ae-slider /deep/ .agile__dots {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 20px;
  margin: 0;
  padding: 0;
}

.ae-slider /deep/ .agile__dots > .agile__dot {
  list-style: none;
  border-radius: 50%;
  background: $color-neutral;
  margin: 0 rem(2px);
  width: 10px;
  height: 10px;
  transition: all $base-transition-time;
  cursor: pointer;

  > button {
    background: none;
    height: 100%;
    width: 100%;
    border: none;
  }

  &--current {
    background: $color-primary;
  }
}

.ae-slider.fullscreen {
  height: 100%;
}

.ae-slider.fullscreen /deep/ .agile__list,
.ae-slider.fullscreen /deep/ .agile__track,
.ae-slider.fullscreen /deep/ .agile__slide {
  height: 100%;
}

.ae-slider.fullscreen /deep/ .agile__dots {
  position: absolute;
  right: 0;
  left: 0;
  bottom: 0;
  height: auto;
  padding: rem(20px) rem(32px);
  z-index: 1000;
}
</style>
