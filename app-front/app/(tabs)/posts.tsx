import React, { useEffect, useRef, useState } from "react";
import {
  StyleSheet,
  Animated,
  ActivityIndicator,
  useColorScheme,
  Image,
  RefreshControl,
  NativeSyntheticEvent,
  NativeScrollEvent,
  FlatList,
} from "react-native";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import { PostPanel } from "@/components/PostPanel";
import { Colors } from "@/constants/Colors";
import { useHomeTabHandler } from "@/providers/HomeTabScrollContext";
import { usePostsScreen } from "@/features/post/usePostsScreen";

export default function PostsScreen() {
  const [refreshing, setRefreshing] = useState(false);
  const colorScheme = useColorScheme();
  const scrollY = useRef(new Animated.Value(0)).current;
  const scrollYRef = useRef(0);
  const listRef = useRef<FlatList>(null);
  const HEADER_HEIGHT = 80;

  const icon =
    colorScheme === "light"
      ? require("../../assets/images/icon-green.png")
      : require("../../assets/images/icon-dark.png");

  const { data, isLoading, error, refetch } = usePostsScreen();

  const handleScroll = (event: NativeSyntheticEvent<NativeScrollEvent>) => {
    scrollYRef.current = event.nativeEvent.contentOffset.y;
  };

  const headerTranslateY = scrollY.interpolate({
    inputRange: [0, HEADER_HEIGHT],
    outputRange: [0, -HEADER_HEIGHT],
    extrapolate: "clamp",
  });

  const { setHandler } = useHomeTabHandler();

  useEffect(() => {
    setHandler(() => {
      listRef.current?.scrollToOffset({
        offset: -(HEADER_HEIGHT + 12),
        animated: true,
      });
    });
  }, [setHandler]);

  if (isLoading) {
    return (
      <ThemedView style={styles.container}>
        <ActivityIndicator size="large" />
      </ThemedView>
    );
  }

  if (error) {
    return (
      <ThemedView style={styles.container}>
        <Animated.FlatList
          data={[]}
          renderItem={() => null}
          ListEmptyComponent={
            <ThemedText style={styles.errorText}>
              ポストが取得できませんでした
            </ThemedText>
          }
          refreshControl={
            <RefreshControl
              refreshing={false}
              onRefresh={() => {}}
              progressViewOffset={HEADER_HEIGHT}
              tintColor={colorScheme === "light" ? "black" : "white"}
            />
          }
          contentInset={{ top: HEADER_HEIGHT }}
          contentOffset={{ x: 0, y: -HEADER_HEIGHT }}
        />
      </ThemedView>
    );
  }

  return (
    <ThemedView style={styles.container}>
      <Animated.View
        style={[
          styles.header,
          {
            transform: [{ translateY: headerTranslateY }],
            backgroundColor: Colors[colorScheme ?? "light"].background,
          },
        ]}
      >
        <Image source={icon} style={styles.logo} />
      </Animated.View>

      <Animated.FlatList
        ref={listRef}
        style={{
          backgroundColor: colorScheme === "light" ? "white" : "black",
        }}
        contentInset={{ top: HEADER_HEIGHT + 12 }}
        contentOffset={{ x: 0, y: -(HEADER_HEIGHT + 12) }}
        contentContainerStyle={{ paddingBottom: 75 }}
        data={data}
        keyExtractor={(item) => item.id}
        renderItem={({ item }) => <PostPanel post={item} />}
        refreshControl={
          <RefreshControl
            refreshing={false}
            onRefresh={() => {}}
            tintColor={colorScheme === "light" ? "black" : "white"}
          />
        }
        onScroll={Animated.event(
          [{ nativeEvent: { contentOffset: { y: scrollY } } }],
          {
            useNativeDriver: true,
            listener: handleScroll,
          }
        )}
        scrollEventThrottle={16}
      />
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  header: {
    position: "absolute",
    top: 0,
    left: 0,
    right: 0,
    height: 80,
    zIndex: 10,
    justifyContent: "flex-start",
    alignItems: "center",
  },
  logo: {
    width: 32,
    height: 32,
    marginTop: 40,
    resizeMode: "contain",
  },
  errorText: {
    textAlign: "center",
    marginTop: 100,
    fontSize: 16,
    color: "gray",
  },
});
