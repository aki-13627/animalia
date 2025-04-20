import React from "react";
import { View, TouchableOpacity, Text, StyleSheet } from "react-native";
import { useColorScheme } from "react-native";
import { Colors } from "@/constants/Colors";

export type ProfileTabType = "posts" | "mypet";

type ProfileTabSelectorProps = {
  selectedTab: ProfileTabType;
  onSelectTab: (tab: ProfileTabType) => void;
};

export const ProfileTabSelector: React.FC<ProfileTabSelectorProps> = ({
  selectedTab,
  onSelectTab,
}) => {
  const colorScheme = useColorScheme();
  const colors = Colors[colorScheme ?? "light"];
  const backgroundColor = colorScheme === "light" ? "white" : "black";
  const styles = getStyles(colors);

  return (
    <View style={[styles.tabContainer, { backgroundColor }]}>
      <TouchableOpacity
        onPress={() => onSelectTab("posts")}
        style={[
          styles.tabButton,
          selectedTab === "posts" && styles.tabButtonActive,
        ]}
      >
        <Text style={styles.tabText}>投稿一覧</Text>
      </TouchableOpacity>
      <TouchableOpacity
        onPress={() => onSelectTab("mypet")}
        style={[
          styles.tabButton,
          selectedTab === "mypet" && styles.tabButtonActive,
        ]}
      >
        <Text style={styles.tabText}>マイペット</Text>
      </TouchableOpacity>
    </View>
  );
};

const getStyles = (colors: typeof Colors.light) =>
  StyleSheet.create({
    tabContainer: {
      flexDirection: "row",
      justifyContent: "center",
      borderColor: colors.icon,
    },
    tabButton: {
      width: "50%",
      paddingVertical: 15,
      paddingHorizontal: 20,
      alignItems: "center",
    },
    tabButtonActive: {
      borderBottomWidth: 2,
      borderColor: colors.tint,
    },
    tabText: {
      fontSize: 16,
      fontWeight: "bold",
      color: colors.text,
    },
  });
