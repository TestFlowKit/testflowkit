import { ref, provide, type InjectionKey, type Ref } from "vue";

export type UseSidebar = {
  isOpen: Ref<boolean>;
  toggleSidebar: () => void;
};

export const useSidebarKey = Symbol() as InjectionKey<UseSidebar>;

export default function useSidebar() {
  const isOpen = ref(true);

  const breakpoint = 1024;
  function toggleSidebar() {
    isOpen.value = !isOpen.value;
  }

  onMounted(() => {
    handleResize();
    window.addEventListener("resize", handleResize);
  });

  onUnmounted(() => {
    window.removeEventListener("resize", handleResize);
  });


  function handleResize() {
    const isDesktop = window.innerWidth >= breakpoint;
    isOpen.value = isDesktop;
  };

  provide(useSidebarKey, {
    isOpen,
    toggleSidebar,
  });
}
